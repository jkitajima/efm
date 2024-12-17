package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jkitajima/efm/lib/composer"
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
	stdin io.Reader,
	stdout io.Writer,
	strerr io.Writer,
	getenv func(string) string,
	getwd func() (string, error),
) error {

	fmt.Println("efm svc api")

	cfg, err := NewConfig(args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db, err := initDB(&cfg.DB)
	if err != nil {
		return err
	}

	inputValidator := validator.New(validator.WithRequiredStructEnabled())

	srv := composer.NewComposer()

	userServer := UserServer.NewServer(db, inputValidator)
	srv.Compose(userServer)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), srv))
	return nil
}

func initDB(config *DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
		config.SSL,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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

func run() {
	// start server
	// stop server (graceful shutdown)
	panic("unimplemented")
}
