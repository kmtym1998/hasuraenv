package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:          "hasuraenv",
		Short:        "Manage multiple hasura-cli versions",
		Long:         "Manage multiple hasura-cli versions. Run 'hasuraenv --help' for usage",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi")
		},
	}
}
