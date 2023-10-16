package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(&printCmd{}, "main")
	subcommands.Register(&writeCmd{}, "main")
	subcommands.Register(subcommands.Alias("p", &printCmd{}), "main")
	subcommands.Register(subcommands.Alias("w", &writeCmd{}), "main")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
