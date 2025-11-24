package config

import (
	"log"
	"github.com/spf13/viper"
)

// Config structure plate pour simplifier le TP
type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
}

// LoadConfig lit la configuration
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// Ce n'est pas grave si le fichier n'est pas trouvé tant qu'on a des variables d'env
		log.Println("Fichier de config non trouvé, utilisation des variables d'environnement ou défauts")
	}

	err = viper.Unmarshal(&config)
	return
}