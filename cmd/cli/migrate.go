package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	// 👇 Toujours le même import vital
	"github.com/axellelanca/urlshortener/cmd"
)

// migrateCmd représente la commande "migrate"
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Initialise la base de données",
	Long:  `Crée ou met à jour les tables (Links, Clicks) dans la base de données SQLite.`,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("🔄 Démarrage des migrations GORM...")

		// TODO: Appeler le modèle (Poste 1) : models.InitDB() ou models.Migrate()
		
		// Simulation
		fmt.Println("✅ Tables 'links' et 'clicks' créées ou mises à jour avec succès (Simulation).")
	},
}

func init() {
	// Pas de drapeau spécifique pour l'instant.
	
	// Enregistrement auprès du Chef (RootCmd)
	cmd.RootCmd.AddCommand(migrateCmd)
}