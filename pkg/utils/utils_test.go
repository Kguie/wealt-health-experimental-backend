package utils

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetIdFromURL_Success(t *testing.T) {
	// Création d'un UUID valide
	validID := uuid.New().String()

	// Création d'une requête factice avec un ID dans l'URL
	req := httptest.NewRequest("GET", "/employees/"+validID, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", validID)

	// Utilisation de `context.WithValue` au lieu de `http.WithValue`
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	// Appel de la fonction à tester
	parsedID, err := GetIdFromURL(req)

	// Vérification des résultats
	assert.NoError(t, err)
	assert.NotNil(t, parsedID)
	assert.Equal(t, validID, parsedID.String())
}

func TestGetIdFromURL_MissingID(t *testing.T) {
	// Création d'une requête sans ID
	req := httptest.NewRequest("GET", "/employees/", nil)
	rctx := chi.NewRouteContext()

	// Utilisation de `context.WithValue`
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	// Appel de la fonction à tester
	parsedID, err := GetIdFromURL(req)

	// Vérification des résultats
	assert.Error(t, err)
	assert.Nil(t, parsedID)
	assert.Contains(t, err.Error(), "ID non trouvé")
}

func TestGetIdFromURL_InvalidUUID(t *testing.T) {
	// Création d'une requête avec un ID invalide
	invalidID := "invalid-uuid"

	req := httptest.NewRequest("GET", "/employees/"+invalidID, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", invalidID)

	// Utilisation de `context.WithValue`
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	req = req.WithContext(ctx)

	// Appel de la fonction à tester
	parsedID, err := GetIdFromURL(req)

	// Vérification des résultats
	assert.Error(t, err)
	assert.Nil(t, parsedID)
	assert.Contains(t, err.Error(), "ID invalide")
}
