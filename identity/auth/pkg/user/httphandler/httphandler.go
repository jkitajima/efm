package httphandler

import (
	"net/http"

	repo "github.com/jkitajima/efm/identity/auth/pkg/user/repo/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/jkitajima/efm/identity/auth/pkg/user"
	"github.com/jkitajima/efm/lib/composer"
	"gorm.io/gorm"
)

type UserServer struct {
	entity  string
	mux     *chi.Mux
	prefix  string
	service *user.Service
	db      user.Repoer
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

func NewServer(db *gorm.DB) composer.Server {
	s := &UserServer{
		entity: "users",
		prefix: "/users",
		mux:    chi.NewRouter(),
		db:     repo.NewRepo(db),
	}
	s.service = &user.Service{Repo: s.db}
	s.addRoutes()
	return s
}
