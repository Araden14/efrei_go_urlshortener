package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	
	// 👇 On importe le paquet cmd pour accéder à RootCmd
	"github.com/axellelanca/urlshortener/cmd" 
)

var shortCode string // La variable pour stocker le code (ex: XYZ123)

// statsCmd représente la commande "stats"
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Affiche les statistiques d'un lien",
	Long:  `Récupère et affiche le nombre de clics pour un code court donné.`,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("📊 Récupération des statistiques pour le code : %s\n", shortCode)
		
		// TODO: Appeler le service (Poste 1) : service.GetStats(shortCode)
		
		// Simulation
		fmt.Println("--- RÉSULTAT (Simulation) ---")
		fmt.Println("URL d'origine : https://www.efrei.fr")
		fmt.Println("Nombre de clics : 42")
	},
}

func init() {
	// 1. Configuration du drapeau --code (ou -c)
	statsCmd.Flags().StringVarP(&shortCode, "code", "c", "", "Le code court du lien (ex: XYZ123)")
	statsCmd.MarkFlagRequired("code") // Obligatoire

	// 2. Enregistrement auprès du Chef (RootCmd)
	cmd.RootCmd.AddCommand(statsCmd)
}