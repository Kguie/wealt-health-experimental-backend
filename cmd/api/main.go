package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	db "wealth-health-backend/pkg/db/postgres"
	"wealth-health-backend/pkg/env"
)

// App structure de l'application
type App struct {
	router http.Handler
	server *http.Server
}

// New initialise l'application et la connexion PostgreSQL
func New() *App {
	// Initialisation de la connexion PostgreSQL
	db.Connect()

	// Charger une seule fois le routeur
	router := loadRouter()

	app := &App{
		router: router,
		server: &http.Server{
			Addr:         env.GetString("ADDR", ":3000"),
			Handler:      router,
			WriteTimeout: time.Second * 30,
			ReadTimeout:  time.Second * 10,
			IdleTimeout:  time.Minute,
		},
	}

	return app
}

// Start démarre l'application
func (a *App) Start(ctx context.Context) error {
	fmt.Println("🚀 Serveur démarré sur", a.server.Addr)

	// Démarrer le serveur HTTP
	err := a.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("❌ Erreur serveur : %w", err)
	}

	return nil
}

// Shutdown arrête proprement le serveur HTTP et PostgreSQL
func (a *App) Shutdown(ctx context.Context) error {
	log.Println("🔄 Arrêt du serveur en cours...")

	// Contexte avec timeout pour éviter un arrêt bloqué
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Arrêter le serveur HTTP proprement
	if err := a.server.Shutdown(shutdownCtx); err != nil {
		log.Printf("❌ Erreur lors de l'arrêt du serveur HTTP : %v", err)
		return err
	}

	// Fermeture de PostgreSQL avec timeout
	dbCtx, dbCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer dbCancel()
	db.Shutdown(dbCtx)

	log.Println("✅ Service API arrêté proprement.")
	return nil
}
