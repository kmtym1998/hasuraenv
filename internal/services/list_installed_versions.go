package services

import (
	"os"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/sirupsen/logrus"
)

func ListInstalledVersions(versionsDir string, logger *logrus.Logger) ([]semver.Version, error) {
	dirEntries, err := os.ReadDir(versionsDir)
	if err != nil {
		return nil, err
	}

	var versions []semver.Version
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			continue
		}

		if err := ValidateSemVer(entry.Name()); err != nil {
			if entry.Name() != "default" {
				logger.Warnf("Unexpected version: %s", entry.Name())
			}
			continue
		}

		sv, err := semver.Parse(strings.Replace(entry.Name(), "v", "", 1))
		if err != nil {
			return nil, err
		}

		versions = append(versions, sv)
	}

	return versions, nil
}
