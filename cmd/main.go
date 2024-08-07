package main

import (
	"fmt"

	"github.com/Biryani-Labs/ezeth/cli"
	"github.com/Biryani-Labs/ezeth/config"
)

func init() {
	config.InitilizeConfig()
}

func main() {
	fmt.Println(config.HOME_DIR)
	cli.StartCli()
}
