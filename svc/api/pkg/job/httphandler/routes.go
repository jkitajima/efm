package httphandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *JobServer) addRoutes() {
	// Private routes
	s.mux.Group(func(r chi.Router) {
		// r.Use(oauth.Authorize("", nil))
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("job created"))
		})
	})
}
