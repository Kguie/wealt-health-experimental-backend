package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wealth-health-backend/cmd/api"
)

func main() {
	// Création d'un contexte avec annulation pour gérer l'arrêt propre du service
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Assure que le contexte est bien annulé à la fin

	// Capture des signaux système (CTRL+C, SIGTERM) pour arrêter proprement l'application
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Initialisation de l'application
	app := api.New()
	log.Println("✅ Wealth Health API is starting...")

	// Lancer l'application dans une goroutine pour ne pas bloquer
	go func() {
		if err := app.Start(ctx); err != nil {
			log.Fatalf("❌ Failed to start app: %v", err)
		}
	}()

	// Attente d'un signal d'arrêt
	<-signalChan
	log.Println("🛑 Shutdown signal received. Stopping Wealth Health API...")

	// Appel de la fonction d'annulation du contexte pour arrêter proprement l'application
	cancel()

	// (Optionnel) Attendre quelques secondes pour laisser l'application terminer proprement
	log.Println("⌛ Cleaning up resources...")
	// time.Sleep(2 * time.Second) // Décommente si des processus asynchrones doivent se terminer

	log.Println("✅ Wealth Health API has stopped gracefully.")
}
