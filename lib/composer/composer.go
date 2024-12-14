package composer

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server interface {
	Mux() http.Handler
	Prefix() string
}

type Composer struct {
	servers []Server
	Mux     *chi.Mux
}

func NewComposer() *Composer {
	return &Composer{Mux: chi.NewRouter()}
}

func (c *Composer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Mux.ServeHTTP(w, r)
}

func (c *Composer) Compose(servers ...Server) error {
	if len(c.servers) > 0 {
		return errors.New("composer: composer is already filled with servers")
	}

	for _, s := range servers {
		prefix := s.Prefix()
		if prefix == "" {
			return errors.New("composer: server prefix is empty")
		}

		Mux := s.Mux()
		if Mux == nil {
			return errors.New("composer: server Mux is nil")
		}

		c.Mux.Mount(prefix, Mux)
	}

	return nil
}
