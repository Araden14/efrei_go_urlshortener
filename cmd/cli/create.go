package cli

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// 👇 C'est l'import crucial qui permet de voir "RootCmd"
	"github.com/axellelanca/urlshortener/cmd"
	"github.com/axellelanca/urlshortener/internal/config"
	"github.com/axellelanca/urlshortener/internal/repository"
	"github.com/axellelanca/urlshortener/internal/services"
)

var urlToShorten string

// createCmd représente la commande "create"
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Crée une URL courte à partir d'une URL longue.",
	Long: `Cette commande raccourcit une URL longue fournie et affiche le code court généré.

Exemple:
  url-shortener create --url="https://www.google.com/search?q=go+lang"`,
	Run: func(cmd *cobra.Command, args []string) {
		if urlToShorten == "" {
			fmt.Println("Erreur: le drapeau --url est obligatoire.")
			os.Exit(1);
		}

		//Vérification de la syntaxe de l'url
		if _, err := url.ParseRequestURI(urlToShorten); err != nil {
			fmt.Println("Erreur: l'url n'est pas valide.");
			os.Exit(1);
		}

		//Charger la configuration
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
		dsn := cfg.Database.Name

		//Connexion à la base de données
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

		//Creation de la short url
		link, err := linkService.CreateLink(urlToShorten);
		if err != nil {
			log.Fatal(err)
		}

		fullShortURL := fmt.Sprintf("%s/%s", cfg.Server.BaseURL, link.Shortcode)
		fmt.Printf("URL courte créée avec succès:\n")
		fmt.Printf("Code: %s\n", link.Shortcode)
		fmt.Printf("URL complète: %s\n", fullShortURL)
	},
}

func init() {
	// 1. On configure les drapeaux (flags) de CETTE commande
	CreateCmd.Flags().StringVarP(&urlToShorten, "url", "u", "", "L'URL à raccourcir (ex: https://google.com)")
	CreateCmd.MarkFlagRequired("url")

	// 2. On attache CETTE commande au chef (RootCmd)
	// C'est "l'architecture inversée" dont tu parlais
	cmd.RootCmd.AddCommand(CreateCmd)
}
