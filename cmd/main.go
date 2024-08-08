package main

import (
	"github.com/Biryani-Labs/ethz/cli"
	"github.com/Biryani-Labs/ethz/config"
)

func init() {
	config.InitilizeConfig()
}

func main() {
	cli.Run()
}
