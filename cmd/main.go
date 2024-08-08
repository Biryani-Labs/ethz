package main

import (
	"github.com/Biryani-Labs/ezeth/cli"
	"github.com/Biryani-Labs/ezeth/common/logs"
	"github.com/Biryani-Labs/ezeth/config"
)

func init() {
	logs.Initilize()
	config.InitilizeConfig()
}

func main() {
	cli.Run()
}
