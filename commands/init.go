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
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := os.MkdirAll(ec.GlobalConfig.HasuraenvPath.VersionsDir+"/latest", os.ModePerm); err != nil {
				return err
			}
			ec.Logger.Info("✅", ec.GlobalConfig.HasuraenvPath.VersionsDir, " has been created")

			ec.Spin("Installing latest hasura cli... ")
			if err := services.InstallHasuraCLI(services.InstallHasuraClIOptions{
				Dir: ec.GlobalConfig.HasuraenvPath.VersionsDir + "/latest",
			}); err != nil {
				return err
			}
			ec.Spinner.Stop()
			ec.Logger.Info("✅", "Latest hasura cli has been installed")

			if err := replaceSymlink(
				ec.GlobalConfig.HasuraenvPath.VersionsDir+"/latest",
				ec.GlobalConfig.HasuraenvPath.Current,
			); err != nil {
				return err
			}

			// FIXME: もとの hasura あればそれは消してねのメッセージ出す
			ec.Logger.InfoFn(func() []interface{} {
				return []interface{}{
					"✅ hasuraenv has been initialized!\n",
					fmt.Sprintf("Run:\n     export PATH=%s:$PATH\n", ec.GlobalConfig.HasuraenvPath.Current),
				}
			})

			return nil
		},
	}
}

func replaceSymlink(dest, src string) error {
	os.Remove(src)
	return os.Symlink(dest, src)
}
