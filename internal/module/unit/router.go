package unit

import (
	"bobobox/internal/repository/postgresql/unit"
	"bobobox/pkg/app"

	"github.com/go-chi/chi"
)

func SetupUnitRoutes(r chi.Router, app app.AppConfig) {
	// Initialize repositories and services
	unitRepository := unit.NewUnitRepository(app)
	unitService := NewUnitService(app, unitRepository)
	unitHandler := NewUnitHandler(app, unitService)

	r.Route("/units", func(r chi.Router) {
		r.Get("/", unitHandler.ListPaginate)
		r.Get("/{id:[a-fA-F0-9-]{36}}", unitHandler.Detail)
		r.Post("/", unitHandler.Create)
		r.Put("/{id:[a-fA-F0-9-]{36}}", unitHandler.Update)
	})
}
