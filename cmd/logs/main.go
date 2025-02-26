package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// basicHandler est la fonction qui gère toutes les requêtes HTTP reçues par le serveur.
func basicHandler(w http.ResponseWriter, r *http.Request) {
	// Écriture de la réponse HTTP. Le contenu est un simple tableau d'octets représentant la chaîne.
	_, err := w.Write([]byte("OK"))
	if err != nil {
		// Si une erreur survient lors de l'écriture, elle est loguée.
		log.Printf("Erreur lors de l'écriture de la réponse : %v", err)
	}
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", basicHandler)
	// Création d'une instance de serveur HTTP.
	// Le serveur écoute sur le port 8080 et utilise basicHandler pour traiter les requêtes.
	server := &http.Server{
		Addr:    ":8080", // Adresse et port d'écoute.
		Handler: router,  // Convertit la fonction basicHandler en un type qui implémente http.Handler.
	}

	// Log indiquant que le serveur démarre et sur quel port il sera à l'écoute.
	log.Printf("Serveur démarré avec succès sur le port %s", server.Addr)

	// Démarrage du serveur en mode bloquant.
	// ListenAndServe démarre la boucle d'écoute et de traitement des requêtes HTTP.
	err := server.ListenAndServe()
	if err != nil {
		// En cas d'erreur critique (par exemple, le port est déjà utilisé), on log l'erreur et on arrête le programme.
		log.Fatal("Failed to start server: ", err)
	}
}
