package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config est la structure qui va stocker notre "menu" (les paramètres)
// C'est comme un bon de commande vide qu'on va remplir.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Monitor  MonitorConfig
}

type ServerConfig struct {
	Port string // Stocke le port (ex: "8080")
}

type DatabaseConfig struct {
	FilePath string // Stocke le chemin du fichier DB
}

type MonitorConfig struct {
	IntervalSeconds int // Stocke le temps entre deux vérifications
}

// LoadConfig lit le fichier config.yaml et remplit la structure Config
func LoadConfig() (*Config, error) {
	// On dit à Viper : "Le fichier s'appelle config"
	viper.SetConfigName("config")
	// On dit à Viper : "C'est un fichier de type yaml"
	viper.SetConfigType("yaml")
	// On dit à Viper : "Cherche ce fichier dans le dossier configs/"
	viper.AddConfigPath("./configs")
	// On ajoute aussi le dossier courant au cas où (.)
	viper.AddConfigPath(".")

	// On essaie de lire le fichier
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Erreur lors de la lecture du fichier de config : %v", err)
		return nil, err
	}

	// On crée une variable vide 'cfg'
	var cfg Config
	// On demande à Viper de verser les infos du fichier dans notre variable 'cfg'
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// On retourne la configuration prête à l'emploi !
	return &cfg, nil
}