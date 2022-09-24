package commands

import (
	"fmt"

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
			versions, err := services.ListInstalledVersions(
				ec.GlobalConfig.HasuraenvPath.VersionsDir,
				ec.Logger,
			)
			if err != nil {
				return err
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
