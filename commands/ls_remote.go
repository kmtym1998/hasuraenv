package commands

import (
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/model"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLsRemoteCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:          "ls-remote",
		Short:        "List remote versions",
		Long:         "List remote versions",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = viper.New()
			return ec.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			logger.SetOutput(ec.Stdout)
			logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, DisableColors: ec.NoColor})
			if !ec.IsTerminal {
				logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
			}

			releases, err := services.ListHasuraReleases()
			if err != nil {
				return err
			}

			logger.InfoFn(func() []interface{} {
				// FIXME: 全部のリリース取得する できたら limit オブション渡せるようにする
				messages := []interface{}{"Latest 30 releases"}
				versions := lo.Map(releases, func(r model.GitHubRelease, _ int) interface{} {
					return "\n     " + r.TagName
				})

				return append(messages, versions...)
			})

			return nil
		},
	}
}
