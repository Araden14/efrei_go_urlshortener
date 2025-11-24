package cmd

import (
	"fmt"
	"log"
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
	Short: "Un service de raccourcissement d'URLs avec API REST et CLI",
	Long: `
'url-shortener' est une application complète pour gérer des URLs courtes.
Elle inclut un serveur API pour le raccourcissement et la redirection,
ainsi qu'une interface en ligne de commande pour l'administration.

Utilisez 'url-shortener [command] --help' pour plus d'informations sur une commande.`,
	// Run: On ne met rien ici, car la commande racine seule affiche juste l'aide par défaut.
}

// Execute est le point d'entrée principal pour l'application Cobra.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init() est une fonction spéciale de Go qui s'exécute automatiquement
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
