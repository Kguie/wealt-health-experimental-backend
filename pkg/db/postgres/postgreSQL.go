package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"wealth-health-backend/ent"
	"wealth-health-backend/pkg/env"

	"entgo.io/ent/dialect"
	entSQL "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

var Client *ent.Client

// Connect initialise la connexion à PostgreSQL avec vérification et logs améliorés
func Connect() {
	// Récupérer les variables d'environnement
	host := env.GetString("POSTGRES_HOST", "")
	user := env.GetString("POSTGRES_USER", "")
	password := env.GetString("POSTGRES_PASSWORD", "")
	dbname := env.GetString("POSTGRES_NAME", "")
	port := env.GetString("POSTGRES_PORT", "")
	sslmode := env.GetString("POSTGRES_SSLMODE", "disable") // Par défaut, désactiver SSL

	// Vérifier si des variables manquent
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("❌ Erreur : Certaines variables d'environnement PostgreSQL sont manquantes. Vérifiez votre configuration.")
	}

	// Construire la chaîne de connexion PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	// Tester la connexion avec `database/sql` avant d'utiliser `Ent`
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Impossible d'ouvrir la connexion à PostgreSQL : %v", err)
	}

	// Vérifier si PostgreSQL répond avec `Ping()`
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Impossible de se connecter à PostgreSQL : %v", err)
	}
	fmt.Println("✅ Connexion établie avec succès avec PostgreSQL !")

	// Ouvrir la connexion avec `Ent`
	drv := entSQL.OpenDB(dialect.Postgres, sqlDB)
	Client = ent.NewClient(ent.Driver(drv))

	// Exécuter la migration automatique
	ctx := context.Background()
	if err := Client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ Erreur lors de la migration du schéma : %v", err)
	}

	fmt.Println("✅ Connexion réussie à PostgreSQL avec Ent !")
}

// Shutdown ferme proprement la connexion PostgreSQL
func Shutdown(ctx context.Context) {
	fmt.Println("🔄 Fermeture de la connexion PostgreSQL...")

	// Timeout pour éviter de bloquer l'API si PostgreSQL ne répond pas
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan error)
	go func() {
		done <- Client.Close()
	}()

	select {
	case <-shutdownCtx.Done():
		log.Println("⚠️ Timeout : Impossible de fermer la connexion PostgreSQL dans le délai imparti.")
	case err := <-done:
		if err != nil {
			log.Printf("❌ Erreur lors de la fermeture de PostgreSQL : %v\n", err)
		} else {
			fmt.Println("✅ Connexion PostgreSQL fermée proprement !")
		}
	}
}
