package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	// 👇 On importe le paquet cmd pour accéder à RootCmd
	"github.com/axellelanca/urlshortener/cmd"
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
		if (code == "") {
			fmt.Println("Erreur: le drapeau --code est obligatoire.")
			os.Exit(1);
		}

		// TODO : Charger la configuration chargée globalement via cmd.cfg


		// TODO 3: Initialiser la connexion à la BDD.
		// log.Fatalf si erreur



		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("FATAL: Échec de l'obtention de la base de données SQL sous-jacente: %v", err)
		}


		// TODO S'assurer que la connexion est fermée à la fin de l'exécution de la commande grâce à defer


		// TODO : Initialiser les repositories et services nécessaires NewLinkRepository & NewLinkService
		linkRepo :=
		linkService :=

		// TODO 5: Appeler GetLinkStats pour récupérer le lien et ses statistiques.
		// Attention, la fonction retourne 3 valeurs
		// Pour l'erreur, utilisez gorm.ErrRecordNotFound
		// Si erreur, os.Exit(1)



		fmt.Printf("Statistiques pour le code court: %s\n", link.ShortCode)
		fmt.Printf("URL longue: %s\n", link.LongURL)
		fmt.Printf("Total de clics: %d\n", totalClicks)
	},
}

func init() {
	rootCmd.AddCommand(StatsCmd);

	//Définir le drapeau
	StatsCmd.Flags().StringVarP(&shortCodeFlag, "code", "c", "", "Code de la short url")
	
	//Définir le drapeau code comme obligatoire
	StatsCmd.MarkFlagRequired("code");
}
