package cmd

import (
	"fmt"
	"os"
	"github.com/Araden14/efrei_go_urlshortener/cmd/cli" // Assure-toi que le chemin d'import est bon !
	"github.com/spf13/cobra"
)

// RootCmd est la commande de base (quand on tape juste ./url-shortener)
var RootCmd = &cobra.Command{
	Use:   "url-shortener",
	Short: "Un raccourcisseur d'URL en Go",
	Long:  `Une application complète pour raccourcir les URLs et suivre les clics.`,
}

// Execute est la fonction appelée par le main.go pour tout lancer
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1) // On quitte le programme s'il y a une erreur grave au démarrage
	}
}

func init() {
	// C'est ICI que tu ajoutes tes "assiettes" au menu principal
	// Tu importes les commandes que tu as créées dans le dossier 'cli'
	
	RootCmd.AddCommand(cli.CreateCmd) // Ajoute la commande 'create'
	RootCmd.AddCommand(cli.StatsCmd)  // Ajoute la commande 'stats'
	
	// Plus tard, tu ajouteras :
	// RootCmd.AddCommand(cli.MigrateCmd)
	// RootCmd.AddCommand(server.ServerCmd)
}