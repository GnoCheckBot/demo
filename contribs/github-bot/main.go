package main

import (
	"context"
	"flag"
	"github-bot/internal/check"
	"github-bot/internal/matrix"
	"os"

	"github.com/gnolang/gno/tm2/pkg/commands"
)

type rootFlags struct {
	verbose bool
}

func main() {
	flags := &rootFlags{}

	cmd := commands.NewCommand(
		commands.Metadata{
			ShortUsage: "github-bot <subcommand> [flags]",
			LongHelp:   "Bot that allows for advanced management of GitHub pull requests.",
		},
		flags,
		commands.HelpExec,
	)

	cmd.AddSubCommands(
		check.NewCheckCmd(flags.verbose),
		matrix.NewMatrixCmd(flags.verbose),
	)

	cmd.Execute(context.Background(), os.Args[1:])
}

func (flags *rootFlags) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(
		&flags.verbose,
		"verbose",
		false,
		"set logging level to debug",
	)
}
