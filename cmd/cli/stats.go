package cli

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/axellelanca/urlshortener/cmd"
	"github.com/axellelanca/urlshortener/internal/config"
	"github.com/axellelanca/urlshortener/internal/repository"
	"github.com/axellelanca/urlshortener/internal/services"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var shortCodeFlag string;// La variable pour stocker le code (ex: XYZ123)

// StatsCmd représente la commande 'stats'
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Affiche les statistiques (nombre de clics) pour un lien court.",
	Long: `Cette commande permet de récupérer et d'afficher le nombre total de clics
pour une URL courte spécifique en utilisant son code.

Exemple:
  url-shortener stats --code="xyz123"`,
	Run: func(cmd *cobra.Command, args []string) {
		if (shortCodeFlag == "") {
			fmt.Println("Erreur: le drapeau --code est obligatoire.")
			os.Exit(1);
		}

		cfg, err := config.LoadConfig()
		dsn := cfg.Database.Name
        db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatalf("FATAL: Échec de l'obtention de la base de données SQL sous-jacente: %v", err)
        }
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		defer sqlDB.Close()

		linkRepo := repository.NewLinkRepository(db);
		linkService := services.NewLinkService(linkRepo);

		//On récupère les statistique du lien via le code
		link, totalClicks, err := linkService.GetLinkStats(shortCodeFlag)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Aucun lien trouvé pour le code %s\n", shortCodeFlag)
			os.Exit(1)
		}
		if err != nil {
			log.Fatalf("FATAL: impossible de récupérer les stats: %v", err)
		}

		fmt.Printf("Statistiques pour le code court: %s\n", link.Shortcode)
		fmt.Printf("URL longue: %s\n", link.LongURL)
		fmt.Printf("Total de clics: %d\n", totalClicks)
	},
}

func init() {
	cmd.RootCmd.AddCommand(StatsCmd);

	//Définir le drapeau
	StatsCmd.Flags().StringVarP(&shortCodeFlag, "code", "c", "", "Code de la short url")
	
	//Définir le drapeau code comme obligatoire
	StatsCmd.MarkFlagRequired("code");
}
