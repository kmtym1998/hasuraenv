package commands

import (
	"fmt"

	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/model"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLsRemoteCmd(ec *cli.ExecutionContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "ls-remote",
		Short:         "List remote versions",
		Long:          "List remote versions",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = viper.New()
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			limit, err := cmd.Flags().GetInt("limit")
			if err != nil {
				return err
			}

			releases, err := services.ListHasuraReleases(limit)
			if err != nil {
				return err
			}

			ec.Logger.InfoFn(func() []interface{} {
				messages := []interface{}{fmt.Sprintf("Latest %d releases", limit)}
				versions := lo.Map(releases, func(r model.GitHubRelease, _ int) interface{} {
					return "\n     " + r.TagName
				})

				return append(messages, versions...)
			})

			return nil
		},
	}

	cmd.Flags().Int("limit", 30, "Maximum number of releases to fetch")

	return cmd
}
