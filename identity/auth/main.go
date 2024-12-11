package main

import (
	"fmt"
	"log"
	"net/http"

	UserServer "github.com/jkitajima/efm/identity/auth/pkg/user/httphandler"
	"github.com/jkitajima/efm/lib/composer"

	repo "github.com/jkitajima/efm/identity/auth/pkg/user/repo/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello auth")

	dsn := "host=127.1.1.1 user=rootuser_identity_auth_local password=passwd_identity_auth_local dbname=identity_auth port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// UUID support for PostgreSQL
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Migrate the schema
	db.AutoMigrate(&repo.UserModel{})

	srv := composer.NewComposer()

	userServer := UserServer.NewServer(db)
	srv.Compose(userServer)
	log.Fatalln(http.ListenAndServe("127.1.1.2:8080", srv))
}
