package services

import (
	"strings"

	"github.com/blang/semver/v4"
)

func ValidateSemVer(s string) error {
	sv := strings.Replace(s, "v", "", 1)
	if _, err := semver.Make(sv); err != nil {
		return err
	}

	return nil
}
