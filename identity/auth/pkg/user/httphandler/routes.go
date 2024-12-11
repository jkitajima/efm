package httphandler

import (
	"github.com/go-chi/chi/v5"
)

func (s *UserServer) addRoutes() {
	s.mux.Group(func(r chi.Router) {
		// r.Post("/", s.handleUserCreate())
		// r.Get("/{fileID}", s.handleFileFindByID())
		// r.Patch("/{fileID}", s.handleFileUpdate())
		// r.Delete("/{fileID}", s.handleFileDelete())

		r.Post("/register", s.handleUserRegister())
		// r.Post("/validate_email", s.handleValidateEmailService())
	})
}
