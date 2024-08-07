package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Biryani-Labs/ezeth/constants"
	"github.com/spf13/viper"
)

func ImportEnv() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(fmt.Errorf("error getting home directory: %s", err))
	}
	configDir := filepath.Join(homeDir, ".ezeth")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("HOME_DIR", configDir)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Panicln(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	for _, element := range constants.ENV {
		if viper.GetString(element) == "" {
			log.Panicln(fmt.Errorf("env variables not present %s", element))
		}
	}
}
