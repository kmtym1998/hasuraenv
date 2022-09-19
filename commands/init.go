package commands

import (
	"fmt"
	"os"

	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/spf13/cobra"
)

func NewInitCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:          "init",
		Short:        "initialize hasuraenv",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO: すでに initialize されてそうかチェック
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.MkdirAll(ec.GlobalConfig.HasuraenvPath.VersionsDir+"/default", os.ModePerm); err != nil {
				return err
			}
			ec.Logger.Info("✅ ", ec.GlobalConfig.HasuraenvPath.VersionsDir, " has been created")

			ec.Spin("Installing latest hasura cli... ")
			if err := services.InstallHasuraCLI(services.InstallHasuraClIOptions{
				Dir: ec.GlobalConfig.HasuraenvPath.VersionsDir + "/default",
			}); err != nil {
				return err
			}
			ec.Spinner.Stop()
			ec.Logger.Info("✅ ", "Latest hasura cli has been installed")

			if err := services.ReplaceSymlink(
				ec.GlobalConfig.HasuraenvPath.VersionsDir+"/default",
				ec.GlobalConfig.HasuraenvPath.Current,
			); err != nil {
				return err
			}

			ec.Logger.InfoFn(func() []interface{} {
				return []interface{}{
					"✅ hasuraenv has been initialized!\n",
					fmt.Sprintf("Run:\n     export PATH=%s:$PATH;\n     which hasura;\n", ec.GlobalConfig.HasuraenvPath.Current),
					fmt.Sprintf("And check if your hasura command executes '%s/default/hasura'", ec.GlobalConfig.HasuraenvPath.Current),
				}
			})

			return nil
		},
	}
}
