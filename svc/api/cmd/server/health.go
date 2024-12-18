package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/go-chi/chi"
	healthPsql "github.com/hellofresh/health-go/v5/checks/postgres"
	"github.com/jkitajima/efm/lib/composer"
)

type HealthServer struct {
	mux    *chi.Mux
	prefix string
}

func (h *HealthServer) Prefix() string {
	return h.prefix
}

func (h *HealthServer) Mux() http.Handler {
	return h.mux
}

func (h *HealthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func setupHealthCheck(config *Config) composer.Server {
	checker := health.NewChecker(
		health.WithCacheDuration(5*time.Second),
		health.WithTimeout(15*time.Second),
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "db",
			Check: healthPsql.New(healthPsql.Config{
				DSN: config.DB.DSN,
			}),
			MaxContiguousFails: 3,
		}),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Printf("health status changed to %q", state.Status)
		}),
	)

	healthServer := &HealthServer{
		mux:    chi.NewMux(),
		prefix: "/healthz",
	}
	healthServer.mux.HandleFunc("/readiness", health.NewHandler(checker))

	return healthServer
}
