package cmd

import (
	"fmt"
	"os"

	"github.com/axellelanca/urlshortener/internal/config"
	"github.com/spf13/cobra"
)

// Cfg est la variable globale qui contiendra la configuration chargée.
var Cfg *config.Config

// RootCmd représente la commande de base.
// IMPORTANT : Elle commence par une Majuscule (RootCmd) pour être visible
// depuis les autres fichiers (server.go, create.go, etc.) qui vont venir s'y accrocher.
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
	// On demande à Cobra d'initialiser la config avant toute commande
	cobra.OnInitialize(initConfig)

	// Comme indiqué dans tes commentaires :
	// On n'ajoute PAS les commandes ici (pas de AddCommand).
	// Ce sont les sous-fichiers qui feront : cmd.RootCmd.AddCommand(...)
}

// initConfig charge la configuration de l'application.
func initConfig() {
	var err error

	// On essaie de charger la config.
	// Note: Assure-toi que la fonction LoadConfig existe bien dans internal/config/config.go
	// et qu'elle ne demande pas d'arguments, ou adapte ici ("." par exemple).
	Cfg, err = config.LoadConfig()
	
	if err != nil {
		// On loggue juste l'avertissement comme demandé
		log.Printf("Attention: Problème lors du chargement de la configuration: %v. Utilisation des valeurs par défaut.", err)
	}
}
