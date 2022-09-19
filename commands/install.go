package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/spf13/cobra"
)

func NewInstallCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:           "install",
		Short:         "Download and install <version>",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Prepare()

			if len(args) == 0 {
				return errors.New("no version specified")
			}

			version := strings.Replace(args[0], "v", "", 1)
			if _, err := semver.Make(version); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ec.Spin(fmt.Sprintf("Installing hasura-cli %s... ", args[0]))
			services.InstallHasuraCLI(services.InstallHasuraClIOptions{
				Dir:     ec.GlobalConfig.HasuraenvPath.VersionsDir + "/" + args[0],
				Version: args[0],
			})

			ec.Spinner.Stop()

			return nil
		},
	}
}
