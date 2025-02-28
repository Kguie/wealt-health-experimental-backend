package employees

import "github.com/go-chi/chi/v5"

func LoadRoutes(router chi.Router) {
	handler := Employee{}

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
	router.Get("/{id}", handler.GetByID)
	router.Put("/{id}", handler.UpdateByID)
	router.Delete("/{id}", handler.DeleteByID)
}
