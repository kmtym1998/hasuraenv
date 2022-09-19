package commands

import (
	"os"

	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLsCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:           "list",
		Aliases:       []string{"ls", "versions"},
		Short:         "List installed versions",
		Long:          "List installed versions",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = viper.New()
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dirEntries, err := os.ReadDir(ec.GlobalConfig.HasuraenvPath.VersionsDir)
			if err != nil {
				return err
			}

			var versions []interface{}
			for _, entry := range dirEntries {
				if !entry.IsDir() {
					continue
				}

				if err := services.ValidateSemVer(entry.Name()); err != nil {
					if entry.Name() != "default" {
						ec.Logger.Warnf("Unexpected version: %s", entry.Name())
					}
					continue
				}

				versions = append(versions, "\n     "+entry.Name())
			}

			ec.Logger.InfoFn(func() []interface{} {
				messages := []interface{}{"Installed hasura cli"}
				return append(messages, versions...)
			})

			return nil
		},
	}
}
