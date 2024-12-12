package main

import (
	"fmt"
	"log"
	"net/http"

	UserServer "github.com/jkitajima/efm/svc/api/pkg/user/httphandler"
	repo "github.com/jkitajima/efm/svc/api/pkg/user/repo/gorm"

	"github.com/jkitajima/efm/lib/composer"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	fmt.Println("efm svc api")

	dsn := "host=127.0.0.1 user=postgres password=passwd dbname=efm port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "api.",
		},
	})

	// UUID support for PostgreSQL
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec("CREATE TYPE user_role AS ENUM('default', 'admin');")

	// Migrate the schema
	db.AutoMigrate(&repo.UserModel{})

	srv := composer.NewComposer()

	userServer := UserServer.NewServer(db)
	srv.Compose(userServer)
	log.Fatalln(http.ListenAndServe("127.2.1.1:8080", srv))
}
