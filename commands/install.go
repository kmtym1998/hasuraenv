package commands

import (
	"errors"
	"fmt"

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

			version := args[0]

			if err := services.ValidateSemVer(version); err != nil {
				return err
			}

			release, err := services.GetReleaseByTagName("hasura", "graphql-engine", version)
			if err != nil {
				return err
			}

			if release == nil {
				return fmt.Errorf("%s does not exist", version)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			version := args[0]

			ec.Spin(fmt.Sprintf("Installing hasura-cli %s... ", version))
			services.InstallHasuraCLI(services.InstallHasuraClIOptions{
				Dir:     ec.GlobalConfig.HasuraenvPath.VersionsDir + "/" + version,
				Version: version,
			})

			ec.Spinner.Stop()

			return nil
		},
	}
}
