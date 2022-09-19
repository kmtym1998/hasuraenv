package services

import (
	"os"
	"os/exec"
)

type InstallHasuraClIOptions struct {
	Dir     string
	Version string
}

// InstallHasura CLI installs hasura-cli
// NOTE: https://hasura.io/docs/latest/hasura-cli/install-hasura-cli/
func InstallHasuraCLI(opt InstallHasuraClIOptions) error {
	installScript, err := exec.Command("curl", "-L", "https://github.com/hasura/graphql-engine/raw/stable/cli/get.sh").Output()
	if err != nil {
		return err
	}

	if err := os.WriteFile("tmp.sh", installScript, os.ModePerm); err != nil {
		return err
	}

	INSTALL_PATH := os.Getenv("INSTALL_PATH")
	VERSION := os.Getenv("VERSION")

	os.Setenv("INSTALL_PATH", opt.Dir)
	os.Setenv("VERSION", opt.Version)

	if _, err := exec.Command("bash", "tmp.sh").Output(); err != nil {
		return err
	}

	os.Setenv("INSTALL_PATH", INSTALL_PATH)
	os.Setenv("VERSION", VERSION)

	if err := os.Remove("tmp.sh"); err != nil {
		return err
	}

	return nil
}
