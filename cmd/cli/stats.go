package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Variable pour stocker le code court demandé
var shortCode string

// StatsCmd définit la commande pour voir les statistiques
var StatsCmd = &cobra.Command{
	Use:   "stats", // La commande : ./url-shortener stats
	Short: "Affiche les statistiques d'un lien",
	Long:  `Affiche le nombre de clics et les infos pour un code court donné.`,

	Run: func(cmd *cobra.Command, args []string) {
		// 1. Vérification que le code est bien là
		if shortCode == "" {
			fmt.Println("Erreur : Veuillez fournir un code avec --code")
			return
		}

		// 2. TEMPORAIRE : Simulation de l'affichage
		fmt.Printf("Recherche des statistiques pour le code : %s...\n", shortCode)
		fmt.Println("TODO: Connecter le service pour récupérer les vrais clics plus tard.")
		
		// Simulation d'un résultat pour voir si ton affichage est joli
		fmt.Println("--- RÉSULTAT FICTIF ---")
		fmt.Printf("Code: %s\n", shortCode)
		fmt.Println("Nombre de clics: 0 (En attente de la base de données)")
	},
}

func init() {
	// On ajoute le drapeau --code (ou -c)
	StatsCmd.Flags().StringVarP(&shortCode, "code", "c", "", "Le code court (ex: XYZ123)")
}