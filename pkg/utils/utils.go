package utils

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetIdFromURL(r *http.Request) (*uuid.UUID, error) {
	// Récupérer le paramètre 'id' depuis l'URL
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		return nil, fmt.Errorf("ID non trouvé")
	}

	// Convertir le paramètre 'id' en UUID
	parsed, err := uuid.Parse(idParam)
	if err != nil {
		return nil, fmt.Errorf("ID invalide: %w", err)
	}
	return &parsed, nil

}
