package main

import (
	"context"
	"github-bot/internal/check"
	"github-bot/internal/matrix"
	"os"

	"github.com/gnolang/gno/tm2/pkg/commands"
)

func main() {
	cmd := commands.NewCommand(
		commands.Metadata{
			ShortUsage: "github-bot <subcommand> [flags]",
			LongHelp:   "Bot that allows for advanced management of GitHub pull requests.",
		},
		commands.NewEmptyConfig(),
		commands.HelpExec,
	)

	cmd.AddSubCommands(
		check.NewCheckCmd(),
		matrix.NewMatrixCmd(),
	)

	cmd.Execute(context.Background(), os.Args[1:])
}
