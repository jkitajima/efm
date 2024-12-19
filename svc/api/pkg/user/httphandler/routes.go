package httphandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jkitajima/efm/lib/responder"
)

func (s *UserServer) addRoutes() {
	// Private routes
	s.mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(s.auth))
		r.Use(responder.RespondAuth(s.auth))

		r.Post("/", s.handleUserCreate())
		r.Get("/{userID}", s.handleUserFindByID())
		r.Patch("/{userID}", s.handleUserUpdateByID())
		r.Delete("/{userID}", s.handleUserSoftDeleteByID())
	})

	// Public routes
	// s.mux.Group(func(r chi.Router) {})
}
