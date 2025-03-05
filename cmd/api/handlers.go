package api

import (
	"net/http"
	"time"
	"wealth-health-backend/internal/employees"
	db "wealth-health-backend/pkg/db/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	//Toutes les routes de l'api seront ici
	router.Route("/v1", func(route chi.Router) {
		route.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		// Ajout des routes `/v1/employees`
		route.Route("/employees", func(r chi.Router) {
			employees.LoadRoutes(r, db.Client)
		})

	})

	return router
}
