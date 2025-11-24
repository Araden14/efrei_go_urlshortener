package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// TODO Créer Config qui est la structure principale qui mappe l'intégralité de la configuration de l'application.
// Les tags `mapstructure` sont utilisés par Viper pour mapper les clés du fichier de config
// (ou des variables d'environnement) aux champs de la structure Go.

type Server struct {
	Port    int
	BaseURL string
}

type Database struct {
	Name string
}

type Analytics struct {
	BufferSize  int
	WorkerCount int
}

type Monitor struct {
	IntervalMinutes int `mapstructure:"interval_minutes"`
}
type Config struct {
	Server
	Database
	Analytics
	Monitor
}

// LoadConfig charge la configuration de l'application en utilisant Viper.
// Elle recherche un fichier 'config.yaml' dans le dossier 'configs/'.
// Elle définit également des valeurs par défaut si le fichier de config est absent ou incomplet.
func LoadConfig() (*Config, error) {
	// TODO Spécifie le chemin où Viper doit chercher les fichiers de config.
	// on cherche dans le dossier 'configs' relatif au répertoire d'exécution.
	var path string = "./configs/"
	viper.AddConfigPath(path)

	// TODO Spécifie le nom du fichier de config (sans l'extension).
	viper.SetConfigName("config")

	// TODO Spécifie le type de fichier de config.
	viper.SetConfigType("yaml")

	// TODO : Définir les valeurs par défaut pour toutes les options de configuration.
	// Ces valeurs seront utilisées si les clés correspondantes ne sont pas trouvées dans le fichier de config
	// ou si le fichier n'existe pas.
	// server.port, server.base_url etc.
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.base_url", "http://localhost:8080")
	viper.SetDefault("database.name", "url_shortener.db")
	viper.SetDefault("analytics.buffer_size", 1000)
	viper.SetDefault("analytics.worker_count", 5)
	viper.SetDefault("monitor.interval_minutes", 5)

	// TODO : Lire le fichier de configuration.
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// TODO 4: Démapper (unmarshal) la configuration lue (ou les valeurs par défaut) dans la structure Config.
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Log  pour vérifier la config chargée
	log.Printf("Configuration loaded: Server Port=%d, DB Name=%s, Analytics Buffer=%d, Monitor Interval=%dmin",
		cfg.Server.Port, cfg.Database.Name, cfg.Analytics.BufferSize, cfg.Monitor.IntervalMinutes)

	return &cfg, nil // Retourne la configuration chargée
}
