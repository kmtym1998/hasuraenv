// Package main is the entrypoint for the command line executable.
package main

import (
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/commands"
	log "github.com/sirupsen/logrus"
)

var (
	version        string = "v0.1.0"
	configPathBase string = "~/.hasuraenv"
)

// main is the entrypoint function
func main() {
	bo := cli.NewBuildOptions(version, configPathBase)
	ec := cli.NewExecutionContext(bo)
	rootCmd := commands.NewRootCmd()

	rootCmd.AddCommand(
		commands.NewVersionCmd(ec),
		commands.NewInitCmd(ec),
		commands.NewLsRemoteCmd(ec),
		commands.NewLsCmd(ec),
		commands.NewInstallCmd(ec),
		commands.NewUseCmd(ec),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
