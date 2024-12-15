package httphandler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jkitajima/efm/lib/composer"
	repo "github.com/jkitajima/efm/svc/api/pkg/user/repo/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/jkitajima/efm/svc/api/pkg/user"
	"gorm.io/gorm"
)

type UserServer struct {
	entity         string
	mux            *chi.Mux
	prefix         string
	service        *user.Service
	db             user.Repoer
	inputValidator *validator.Validate
}

func (s *UserServer) Prefix() string {
	return s.prefix
}

func (s *UserServer) Mux() http.Handler {
	return s.mux
}

func (s *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(db *gorm.DB, validtr *validator.Validate) composer.Server {
	s := &UserServer{
		entity:         "users",
		prefix:         "/users",
		mux:            chi.NewRouter(),
		db:             repo.NewRepo(db),
		inputValidator: validtr,
	}

	s.service = &user.Service{Repo: s.db}
	s.addRoutes()
	return s
}
