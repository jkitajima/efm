package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/go-chi/oauth"
	UserServer "github.com/jkitajima/efm/identity/auth/pkg/user/httphandler"
	"github.com/jkitajima/efm/lib/composer"

	repo "github.com/jkitajima/efm/identity/auth/pkg/user/repo/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello auth")

	dsn := "host=127.0.0.1 user=postgres password=passwd dbname=identity port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// UUID support for PostgreSQL
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Migrate the schema
	db.AutoMigrate(&repo.UserModel{})

	srv := composer.NewComposer()
	srv.Mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	oauth.NewBearerServer("", 120*time.Second, nil, nil)

	userServer := UserServer.NewServer(db)
	srv.Compose(userServer)
	log.Fatalln(http.ListenAndServe("127.1.1.2:8080", srv))
}
