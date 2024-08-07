package config

import (
	"github.com/Biryani-Labs/ezeth/common/utils"
	"github.com/spf13/viper"
)

var (
	HOME_DIR = ""
)

func InitilizeConfig() {
	utils.ImportEnv()
	HOME_DIR = viper.GetString("HOME_DIR")
}
