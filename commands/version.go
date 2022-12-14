package commands

import (
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewVersionCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:           "version",
		Short:         "Print the CLI version",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			ec.Viper = viper.New()

			return ec.Prepare()
		},
		Run: func(cmd *cobra.Command, args []string) {
			logger := logrus.New()
			logger.SetOutput(ec.Stdout)
			logger.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, DisableColors: ec.NoColor})
			if !ec.IsTerminal {
				logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: false})
			}

			logger.WithField("version", ec.Version).Info("hasuraenv")
		},
	}
}
