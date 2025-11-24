package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	// 👇 On importe tous tes dossiers
	"github.com/axellelanca/urlshortener/cmd"
	"github.com/axellelanca/urlshortener/internal/api"
	"github.com/axellelanca/urlshortener/internal/monitor"
	"github.com/axellelanca/urlshortener/internal/repository"
	"github.com/axellelanca/urlshortener/internal/services"
	"github.com/axellelanca/urlshortener/internal/workers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ServerCmd : La commande qui lance tout
var ServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: "Lance le serveur API et les workers",
	Run: func(c *cobra.Command, args []string) {
		// 1. Chargement de la config (depuis root.go)
		// Si cmd.Cfg est nil, on utilise des valeurs par défaut pour éviter le crash
		port := ":8080"
		dbUrl := "url_shortener.db"
		
		if cmd.Cfg != nil {
			port = cmd.Cfg.ServerPort
			dbUrl = cmd.Cfg.DBUrl
		}

		fmt.Printf("🚀 Démarrage du serveur sur le port %s...\n", port)

		// 2. Connexion Base de Données (SQLite)
		db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Impossible de se connecter à la BDD: %v", err)
		}

		// 3. Initialisation des couches (La "Mise en place")
		
		// Repositories (Cuisine)
		linkRepo := repository.NewLinkRepository(db)
		clickRepo := repository.NewClickRepository(db)

		// Services (Chefs de partie)
		// On passe nil pour clickRepo dans LinkService pour l'instant si ton NewLinkService ne prend qu'un argument.
		// Adapte selon ton fichier link_service.go
		linkService := services.NewLinkService(linkRepo, clickRepo) 
		
		// Handlers (Serveurs)
		apiHandler := api.NewHandler(linkService)

		// Worker (Plongeur)
		clickWorker := workers.NewClickWorker()
		clickWorker.Start()

		// Monitor (Inspecteur)
		urlMonitor := monitor.NewURLMonitor(linkService) // Attention à la majuscule URL
		urlMonitor.Start()

		// 4. Configuration du serveur Web (Gin)
		router := gin.Default()
		apiHandler.RegisterRoutes(router)

		// 5. Lancement du serveur (avec arrêt propre)
		srv := &http.Server{
			Addr:    port,
			Handler: router,
		}

		// On lance le serveur dans un "thread" à part (goroutine)
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("❌ Erreur serveur: %s\n", err)
			}
		}()

		// 6. Attente du signal d'arrêt (Ctrl+C)
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit // On bloque ici tant qu'on n'a pas appuyé sur Ctrl+C
		
		fmt.Println("\n🛑 Arrêt du serveur en cours...")
		
		// On laisse 5 secondes au serveur pour finir les requêtes en cours
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Arrêt forcé du serveur:", err)
		}

		fmt.Println("👋 Serveur arrêté proprement.")
	},
}