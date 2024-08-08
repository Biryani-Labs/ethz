package cli

import (
	"fmt"
	"os"

	cmdconfig "github.com/Biryani-Labs/ezeth/cli/cmd_config"
	cmddraft "github.com/Biryani-Labs/ezeth/cli/cmd_draft"
	cmdexec "github.com/Biryani-Labs/ezeth/cli/cmd_exec"
	"github.com/Biryani-Labs/ezeth/common/logs"
	"github.com/alecthomas/kong"
)

type CLI struct {
	Exec   cmdexec.ExecCmd     `cmd:"" help:"Execute blueprint based on the name"`
	Draft  cmddraft.DraftCmd   `cmd:"" help:"Draft the base for creating the blueprint"`
	Config cmdconfig.ConfigCmd `cmd:"" help:"Command to configure the base for creating the blueprint through cli actions"`
}

func Run() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("ezeth"),
		kong.Description("A cmd wizad for setting up ethereum based validators"),
	)

	if err := ctx.Run(&cli); err != nil {
		logAndExit(fmt.Errorf("error running command: %w", err))
	}
}

func logAndExit(err error) {
	logs.Error(err, "Error initializing Ezeth")
	os.Exit(1)
}
