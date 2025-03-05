package employees

import (
	"wealth-health-backend/ent"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(router chi.Router, client *ent.Client) {
	handler := NewEmployeeHandler(client)

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
	router.Get("/{id}", handler.GetByID)
	router.Put("/{id}", handler.UpdateByID)
	router.Delete("/{id}", handler.DeleteByID)
}
