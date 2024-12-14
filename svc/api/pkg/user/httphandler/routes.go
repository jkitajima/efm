package httphandler

import (
	"github.com/go-chi/chi/v5"
)

func (s *UserServer) addRoutes() {
	// Private routes
	s.mux.Group(func(r chi.Router) {
		// r.Use(oauth.Authorize("", nil))
		r.Post("/", s.handleUserCreate())
		r.Get("/{userID}", s.handleUserFindByID())
		r.Patch("/{userID}", s.handleUserUpdateByID())
		r.Delete("/{userID}", s.handleUserSoftDeleteByID())
	})
}
