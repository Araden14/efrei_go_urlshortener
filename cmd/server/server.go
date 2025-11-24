package server

import (
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/axellelanca/urlshortener/internal/config"
	"github.com/axellelanca/urlshortener/internal/monitor"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// Driver SQLite pour GORM
)

var ServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: "Lance le serveur API de raccourcissement d'URLs et les processus de fond.",
	Long: `Cette commande initialise la base de données, configure les APIs,
démarre les workers asynchrones pour les clics et le moniteur d'URLs,
puis lance le serveur HTTP.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO : créer une variable qui stock la configuration chargée globalement via cmd.cfg
		// Ne pas oublier la gestion d'erreur et faire un fatalF
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("FATAL: Échec du chargement de la configuration: %v", err)
		}

		// TODO : Initialiser la connexion à la bBDD
		dsn := cfg.Database.Name
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("FATAL: Échec de la connexion à la base de données: %v", err)
		}

		// TODO : Initialiser les repositories.
		// Créez des instances de GormLinkRepository et GormClickRepository.

		// Laissez le log
		log.Println("Repositories initialisés.")

		// TODO : Initialiser les services métiers.
		// Créez des instances de LinkService et ClickService, en leur passant les repositories nécessaires.

		// Laissez le log
		log.Println("Services métiers initialisés.")

		// TODO : Initialiser le channel ClickEventsChannel (api/handlers) des événements de clic et lancer les workers (StartClickWorkers).
		// Le channel est bufferisé avec la taille configurée.
		// Passez le channel et le clickRepo aux workers.

		// TODO : Remplacer les XXX par les bonnes variables
		log.Printf("Channel d'événements de clic initialisé avec un buffer de %d. %d worker(s) de clics démarré(s).",
			XXX, XXX)

		// TODO : Initialiser et lancer le moniteur d'URLs.
		// Utilisez l'intervalle configuré
		monitorInterval := time.Duration(XXX) * time.Minute
		urlMonitor := monitor.NewUrlMonitor() // Le moniteur a besoin du linkRepo et de l'interval

		// TODO Lancez le moniteur dans sa propre goroutine.

		log.Printf("Moniteur d'URLs démarré avec un intervalle de %v.", monitorInterval)

		// TODO : Configurer le routeur Gin et les handlers API.
		// Passez les services nécessaires aux fonctions de configuration des routes.

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