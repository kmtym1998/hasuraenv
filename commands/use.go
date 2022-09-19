package commands

import (
	"errors"

	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/spf13/cobra"
)

func NewUseCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:           "use",
		Short:         "Use <version>",
		Long:          "Use <version>",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Prepare()

			if len(args) == 0 {
				return errors.New("no version specified")
			}

			// TODO: ローカルに存在するバージョンかどうか検証
			// TODO: ローカルに存在しないバージョンだったらインストールするかどうか聞く

			return services.ValidateSemVer(args[0])
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return services.ReplaceSymlink(
				ec.GlobalConfig.HasuraenvPath.VersionsDir+"/"+args[0],
				ec.GlobalConfig.HasuraenvPath.Current,
			)
		},
	}
}
