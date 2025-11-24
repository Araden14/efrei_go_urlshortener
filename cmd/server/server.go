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

	"github.com/axellelanca/urlshortener/cmd"
	"github.com/axellelanca/urlshortener/internal/api"
	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/monitor"
	"github.com/axellelanca/urlshortener/internal/repository"
	"github.com/axellelanca/urlshortener/internal/services"
	"github.com/axellelanca/urlshortener/internal/workers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: "Lance le serveur API et les workers",
	Run: func(c *cobra.Command, args []string) {
		// 1. Config
		port := ":8080"
		dbUrl := "url_shortener.db"
		if cmd.Cfg != nil {
			port = cmd.Cfg.ServerPort
			dbUrl = cmd.Cfg.DBUrl
		}

		// 2. Base de données
		db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Erreur DB: %v", err)
		}
		// Migration automatique (Création des tables si elles n'existent pas)
		db.AutoMigrate(&models.Link{}, &models.Click{})

		// 3. Initialisation des composants
		
		// Repositories
		linkRepo := repository.NewLinkRepository(db)
		clickRepo := repository.NewClickRepository(db)

		// Channel pour les workers (Tampon de 100 événements)
		// C'est le tuyau entre l'API et les Workers
		clickChan := make(chan models.ClickEvent, 100)

		// Services
		// Note : Il faudra peut-être adapter NewLinkService selon ce que ton collègue a écrit.
		// J'assume ici qu'il prend le Repo + le Channel (ou juste le repo, à vérifier dans link_service.go)
		// Pour l'instant, je laisse une version standard :
		linkService := services.NewLinkService(linkRepo)
		// SI ton service a besoin du channel pour envoyer les clics, il faudra modifier link_service.go

		// Workers
		// On lance 3 ouvriers pour gérer les clics en parallèle
		workers.StartClickWorkers(3, clickChan, clickRepo)

		// Monitor
		// Vérifie toutes les 2 minutes
		urlMonitor := monitor.NewUrlMonitor(linkRepo, 2*time.Minute)
		go urlMonitor.Start() // On le lance dans une goroutine pour ne pas bloquer

		// API Handlers
		// On passe le service ET le channel au handler (pour qu'il puisse envoyer des événements)
		apiHandler := api.NewHandler(linkService) 
		// ⚠️ ATTENTION : Il faudra probablement modifier api/handlers.go pour qu'il accepte le channel 'clickChan'
		// ou passer le channel au LinkService. C'est le point délicat de l'intégration.

		// 4. Serveur Web
		router := gin.Default()
		apiHandler.RegisterRoutes(router)

		srv := &http.Server{
			Addr:    port,
			Handler: router,
		}

		go func() {
			fmt.Printf("🚀 Serveur démarré sur %s\n", port)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("❌ Erreur serveur: %s\n", err)
			}
		}()

		// 5. Arrêt propre
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		fmt.Println("\n🛑 Arrêt en cours...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Arrêt forcé:", err)
		}
		fmt.Println("👋 Serveur éteint.")
	},
}