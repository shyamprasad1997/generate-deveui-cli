package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var Config Configuration

// Load - To load configuration file
func Load() {
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
