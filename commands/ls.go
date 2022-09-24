package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver/v4"
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/samber/lo"
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

			var versions []semver.Version
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

				sv, err := semver.Parse(strings.Replace(entry.Name(), "v", "", 1))
				if err != nil {
					return err
				}

				versions = append(versions, sv)
			}

			semver.Sort(versions)

			ec.Logger.InfoFn(func() []interface{} {
				messages := []interface{}{"Installed hasura cli"}
				return append(
					messages,
					lo.Map(versions, func(v semver.Version, _ int) any {
						return fmt.Sprint("\n    " + v.String())
					})...,
				)
			})

			return nil
		},
	}
}
