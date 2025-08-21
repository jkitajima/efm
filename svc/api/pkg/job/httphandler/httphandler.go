package httphandler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jkitajima/efm/lib/composer"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type JobServer struct {
	mux    *chi.Mux
	prefix string
}

func (s *JobServer) Prefix() string {
	return s.prefix
}

func (s *JobServer) Mux() http.Handler {
	return s.mux
}

func (s *JobServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer(db *gorm.DB, validtr *validator.Validate) composer.Server {
	s := &JobServer{
		prefix: "/jobs",
		mux:    chi.NewRouter(),
	}

	s.mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("job created"))
	})
	// s.addRoutes()
	return s
}
