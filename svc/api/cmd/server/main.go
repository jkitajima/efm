package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	serverComposer "github.com/jkitajima/efm/lib/composer"
	UserServer "github.com/jkitajima/efm/svc/api/pkg/user/httphandler"
	repo "github.com/jkitajima/efm/svc/api/pkg/user/repo/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	ctx := context.Background()
	if err := exec(ctx, os.Args, os.Stdin, os.Stdout, os.Stderr, os.Getenv, os.Getwd); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func exec(
	ctx context.Context,
	args []string,
	_ io.Reader,
	_ io.Writer,
	_ io.Writer,
	_ func(string) string,
	_ func() (string, error),
) error {
	cfg, err := NewConfig(args)
	if err != nil {
		return err
	}

	// Setting up dependencies
	jwtAuth := jwtauth.New("HS256", []byte(cfg.Auth.SignKey), nil)

	db, err := initDB(&cfg.DB)
	if err != nil {
		return err
	}

	inputValidator := validator.New(validator.WithRequiredStructEnabled())

	// Mounting routers
	composer := serverComposer.NewComposer()
	healthCheck := SetupHealthCheck(cfg)
	userServer := UserServer.NewServer(jwtAuth, db, inputValidator)
	if err := composer.Compose(healthCheck, userServer); err != nil {
		return err
	}

	// Server config
	server := &http.Server{
		Addr:         net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		WriteTimeout: time.Second * time.Duration(cfg.Server.Timeout.Write),
		ReadTimeout:  time.Second * time.Duration(cfg.Server.Timeout.Read),
		IdleTimeout:  time.Second * time.Duration(cfg.Server.Timeout.Idle),
		Handler:      composer,
	}

	// Graceful shutdown
	notifyCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverChan := make(chan error, 1)
	go func() {
		<-notifyCtx.Done()

		timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(cfg.Server.Timeout.Shutdown))
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			serverChan <- err
		}
		serverChan <- nil
	}()

	log.Printf("server listening on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-serverChan
}

func initDB(config *DB) (*gorm.DB, error) {
	config.DSN = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.SSL,
	)
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "api.",
		},
	})
	if err != nil {
		return &gorm.DB{}, err
	}

	// UUID support for PostgreSQL
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Add "user_role" enum data type
	db.Exec(`
		DO $$
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
					CREATE TYPE user_role AS ENUM ('default', 'admin');
				END IF;
			END
		$$;
	`)

	// Migrate the schema
	db.AutoMigrate(&repo.UserModel{})

	return db, nil
}
