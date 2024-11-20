package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pazams/go-create-api/pkg/api/config"
	"github.com/pazams/go-create-api/pkg/api/controllers"
)

// NewRouter ..
func NewRouter(
	ic *controllers.PingController,
	c *config.Config,
) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	if c.HasProxy {
		r.Use(middleware.RealIP)
	}
	if c.AppEnv != "test" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/ping", toHandler(ic.Ping))
	})

	return r
}
