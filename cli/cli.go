package cli

import (
	"github.com/alecthomas/kong"
	"github.com/pterm/pterm"
)

type CLI struct {
}

func StartCli() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("korg"),
		kong.Description("A cmd app for CI/CD pipelines"),
	)
	if err := ctx.Run(&cli); err != nil {
		pterm.Error.PrintOnError(err)
	}
}
