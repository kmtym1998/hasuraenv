package commands

import (
	"fmt"

	cli "github.com/kmtym1998/hasuraenv"

	"github.com/spf13/cobra"
)

// rootCmd is root command
var rootCmd = &cobra.Command{
	Use:          "hasuraenv",
	Short:        "Manage multiple hasura-cli versions",
	Long:         "Manage multiple hasura-cli versions. Run 'hasuraenv --help' for usage",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi")
	},
}

func init() {
	ec := cli.NewExecutionContext()
	rootCmd.AddCommand(
		NewVersionCmd(ec),
		NewInitCmd(ec),
		NewLsRemoteCmd(ec),
		NewLsCmd(ec),
		NewInstallCmd(ec),
		NewUseCmd(ec),
	)
}

func Execute() error {
	return rootCmd.Execute()
}
