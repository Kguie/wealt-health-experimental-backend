package employees

import (
	"encoding/json"
	"net/http"
	"wealth-health-backend/ent"
	"wealth-health-backend/pkg/utils"

	"github.com/go-chi/chi/v5"
)

func LoadRoutes(router chi.Router, client *ent.Client) {
	handler := NewEmployeeHandler(client)

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Lire et décoder le corps JSON dans la structure CreateEmployeeDTO
		var dto CreateEmployeeDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "❌ Erreur de décodage JSON: "+err.Error(), http.StatusMethodNotAllowed)
			return
		}

		// Valider les données
		if err := ValidateEmployee(dto); err != nil {
			http.Error(w, "❌ Erreur de validation: "+err.Error(), http.StatusBadRequest)
			return
		}

		response, err := handler.Create(ctx, dto)
		if err != nil {
			http.Error(w, "❌ Erreur de validation: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		employees, err := handler.List(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	})

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		parsed, err := utils.GetIdFromURL(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		employee, err := handler.GetByID(ctx, *parsed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)
	})

	router.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		parsed, err := utils.GetIdFromURL(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var dto UpdateEmployeeDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "❌ Erreur de décodage JSON: "+err.Error(), http.StatusMethodNotAllowed)
			return
		}

		if err := ValidateEmployeeUpdate(dto); err != nil {
			http.Error(w, "❌ Erreur de validation: "+err.Error(), http.StatusBadRequest)
			return
		}

		employee, err := handler.UpdateByID(ctx, *parsed, dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

	})

	router.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		parsed, err := utils.GetIdFromURL(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := handler.DeleteByID(ctx, *parsed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}
