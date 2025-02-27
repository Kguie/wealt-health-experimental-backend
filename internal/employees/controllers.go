package employees

import "github.com/go-chi/chi/v5"

func LoadRoutes(router chi.Router) {
	employeeHandler := Employee{}

	router.Post("/", employeeHandler.Create)
	router.Get("/", employeeHandler.List)
	router.Get("/{id}", employeeHandler.GetByID)
	router.Put("/{id}", employeeHandler.UpdateByID)
	router.Delete("/{id}", employeeHandler.DeleteByID)
}
