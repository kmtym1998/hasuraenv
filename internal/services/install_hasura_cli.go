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

	const tmpFileName = "tmp_hasuraenv_install_script.sh"

	if err := os.WriteFile(tmpFileName, installScript, os.ModePerm); err != nil {
		return err
	}

	installPath := os.Getenv("INSTALL_PATH")
	version := os.Getenv("VERSION")

	os.Setenv("INSTALL_PATH", opt.Dir)
	os.Setenv("VERSION", opt.Version)

	if _, err := exec.Command("bash", tmpFileName).Output(); err != nil {
		return err
	}

	os.Setenv("INSTALL_PATH", installPath)
	os.Setenv("VERSION", version)

	if err := os.Remove(tmpFileName); err != nil {
		return err
	}

	return nil
}
