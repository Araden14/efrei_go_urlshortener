package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	
	// 👇 C'est l'import crucial qui permet de voir "RootCmd"
	"github.com/axellelanca/urlshortener/cmd" 
)

var urlToShorten string

// createCmd représente la commande "create"
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Crée une nouvelle URL courte",
	Long:  `Raccourcit une URL longue et retourne le code unique généré.`,
	Run: func(c *cobra.Command, args []string) {
		// Logique de la commande
		if urlToShorten == "" {
			fmt.Println("❌ Erreur: L'URL est obligatoire. Utilisez le flag --url")
			return
		}

		fmt.Printf("📝 Création du lien pour : %s\n", urlToShorten)
		
		// TODO: Appeler le service (Poste 1) plus tard
		fmt.Println("✅ Code généré : XYZ123 (Simulation)")
	},
}

func init() {
	// 1. On configure les drapeaux (flags) de CETTE commande
	createCmd.Flags().StringVarP(&urlToShorten, "url", "u", "", "L'URL à raccourcir (ex: https://google.com)")
	createCmd.MarkFlagRequired("url")

	// 2. On attache CETTE commande au chef (RootCmd)
	// C'est "l'architecture inversée" dont tu parlais
	cmd.RootCmd.AddCommand(createCmd)
}