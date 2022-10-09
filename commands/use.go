package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
	cli "github.com/kmtym1998/hasuraenv"
	"github.com/kmtym1998/hasuraenv/internal/services"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func NewUseCmd(ec *cli.ExecutionContext) *cobra.Command {
	return &cobra.Command{
		Use:           "use",
		Short:         "Use <version>",
		Long:          "Use <version>",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := ec.Prepare(); err != nil {
				return err
			}

			if len(args) == 0 {
				return errors.New("no version specified")
			}

			receivedVersionStr := args[0]

			if err := services.ValidateSemVer(receivedVersionStr); err != nil {
				return err
			}

			receivedVersion, err := semver.Parse(strings.Replace(receivedVersionStr, "v", "", 1))
			if err != nil {
				return err
			}

			installedVersions, err := services.ListInstalledVersions(
				ec.GlobalConfig.HasuraenvPath.VersionsDir,
				ec.Logger,
			)
			if err != nil {
				return err
			}

			if _, found := lo.Find(installedVersions, func(installedVersion semver.Version) bool {
				return installedVersion.String() == receivedVersion.String()
			}); found {
				return nil
			}

			prompt := promptui.Prompt{
				Label: fmt.Sprintf("%s is not installed. Would you like to install? (type 'y' to install)", receivedVersionStr),
			}

			input, err := prompt.Run()
			if err != nil {
				return err
			}

			if strings.ToUpper(input) == "Y" || strings.ToUpper(input) == "YES" {
				installCmd := NewInstallCmd(ec)

				if err := installCmd.PreRunE(cmd, []string{receivedVersionStr}); err != nil {
					return err
				}

				if err := installCmd.RunE(cmd, []string{receivedVersionStr}); err != nil {
					return err
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return services.ReplaceSymlink(
				ec.GlobalConfig.HasuraenvPath.VersionsDir+"/"+args[0],
				ec.GlobalConfig.HasuraenvPath.Current,
			)
		},
	}
}
