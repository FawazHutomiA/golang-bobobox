package router

import (
	"bobobox/internal/module/unit"
	"bobobox/pkg/app"

	"github.com/go-chi/chi"
)

func SetupRoutes(r *chi.Mux, app app.AppConfig) {
	// API V1
	r.Route("/api/v1", func(r chi.Router) {
		unit.SetupUnitRoutes(r, app)
	})
}
