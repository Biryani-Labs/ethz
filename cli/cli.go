package cli

import (
	"log"

	cmdconfig "github.com/Biryani-Labs/ethz/cli/cmd_config"
	cmddraft "github.com/Biryani-Labs/ethz/cli/cmd_draft"
	cmdexec "github.com/Biryani-Labs/ethz/cli/cmd_exec"
	"github.com/Biryani-Labs/ethz/common/logs"
	"github.com/Biryani-Labs/ethz/common/utils"
	"github.com/Biryani-Labs/ethz/config"
	"github.com/alecthomas/kong"
)

type CLI struct {
	Exec    cmdexec.ExecCmd     `cmd:"" help:"Execute blueprint based on the name"`
	Draft   cmddraft.DraftCmd   `cmd:"" help:"Draft the base for creating the blueprint"`
	Config  cmdconfig.ConfigCmd `cmd:"" help:"Command to configure the base for creating the blueprint through cli actions"`
	AskPass bool                `help:"Prompt for password" default:"false"`
	Home    string              `help:"Specify the home directory" default:"~/.ethz"`
}

func Run() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("ethz"),
		kong.Description("A cmd wizad for setting up ethereum based validators"),
		kong.HelpOptions{
			Summary: true,
			Compact: true,
		},
	)

	homevar, err := utils.ExpandPath(cli.Home)
	if err != nil {
		logs.Error(err, "Unable to fetch the Home data")
	}

	config.SetHomevar(homevar)
	config.SetAskpass(cli.AskPass)

	if err := ctx.Run(&cli); err != nil {
		log.Fatalln(err)
	}
}
