package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	UserServer "github.com/jkitajima/efm/identity/auth/pkg/user/httphandler"
	"github.com/jkitajima/efm/lib/composer"

	repo "github.com/jkitajima/efm/identity/auth/pkg/user/repo/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
	srv.Mux.Use(middleware.Recoverer)

	srv.Mux.Post("/oauth/token", exchangeToken)

	tokenAuth := jwtauth.New("HS256", []byte("oi"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	jwtauth.Verifier(tokenAuth)
	jwtauth.Authenticator(tokenAuth)

	userServer := UserServer.NewServer(db)
	srv.Compose(userServer)
	log.Fatalln(http.ListenAndServe("127.1.1.2:8080", srv))
}

func exchangeToken(w http.ResponseWriter, r *http.Request) {
	// Check if "Content-Type" header is "application/x-www-form-urlencoded"
	contents := r.Header["Content-Type"]
	if len(contents) != 1 {
		w.Write([]byte("400: must have only one content-type"))
		return
	} else if contents[0] != "application/x-www-form-urlencoded" {
		w.Write([]byte("400: invalid content-type"))
		return
	}

	if err := r.ParseForm(); err != nil {
		w.Write([]byte("400 bad request"))
		return
	}

	// Validate if the required params
	// ("grant_type", "password" and "username")
	// were sent by the client
	requiredParams := [3]string{"grant_type", "password", "username"}
	for _, p := range requiredParams {
		if !r.Form.Has(p) {
			w.Write([]byte("400 bad request: did not sent required params"))
			return
		}
	}

	// request is valid
	// auth user to decide if acess token must be sent or not...
}
