package config

import (
	"path"

	"github.com/Biryani-Labs/ethz/common/utils"
	"github.com/spf13/viper"
)

var (
	HOME_DIR = ""
)

func LocateInHomePath(location string) string {
	return path.Join(HOME_DIR, location)
}

func InitilizeConfig() {
	utils.ImportEnv()
	HOME_DIR = viper.GetString("HOME_DIR")
}
