package commands

import (
	"fmt"
	"os/exec"

	cli "github.com/kmtym1998/hasuraenv"
	"github.com/spf13/cobra"
)

func NewInstallCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:          "install",
		Short:        "Download and install <version>",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// FIXME: args のバリデーション
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			installCmd := fmt.Sprintf(
				"curl -L https://github.com/hasura/graphql-engine/raw/stable/cli/get.sh | INSTALL_PATH=%s VERSION=%s bash",
				ec.GlobalConfig.HasuraenvPath.VersionsDir,
				args[0],
			)

			ec.Logger.Info(installCmd)

			if err := exec.Command(installCmd).Run(); err != nil {
				return err
			}

			return nil
		},
	}
}
