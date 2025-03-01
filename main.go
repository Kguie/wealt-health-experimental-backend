package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wealth-health-backend/cmd/api"
)

func main() {
	// Création d'un contexte avec annulation
	ctx, cancel := context.WithCancel(context.Background())

	// Capture des signaux système (CTRL+C, SIGTERM)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Initialisation et démarrage de l'API
	app := api.New()
	log.Println("✅ Wealth Health API is starting...")

	// Lancer l'API dans une goroutine
	go func() {
		if err := app.Start(ctx); err != nil {
			log.Fatalf("❌ Erreur lors du démarrage de l'API : %v", err)
		}
	}()

	// Attente d'un signal d'arrêt
	<-signalChan
	log.Println("🛑 Signal de fermeture reçu pour l'API.")

	// Annuler le contexte (fermeture propre des services)
	cancel()

	// Arrêt propre de l'API
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("❌ Erreur lors du shutdown : %v", err)
	}

	log.Println("✅ Wealth Health API s'est arrêtée proprement.")
}
