package api

import (
	"net/http"
	employeesService "wealth-health-backend/internal/employees"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/employees", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	employeeHandler := &employeesService.Employee{}

	router.Post("/", employeeHandler.Create)
	router.Get("/", employeeHandler.List)
	router.Get("/{id}", employeeHandler.GetByID)
	router.Put("/{id}", employeeHandler.UpdateByID)
	router.Delete("/{id}", employeeHandler.DeleteByID)
}
