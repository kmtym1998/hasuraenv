package e2e

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// NOTE: https://hasura.io/docs/latest/guides/telemetry/#disable-cli-telemetry
	os.Setenv("HASURA_GRAPHQL_ENABLE_TELEMETRY", "false")

	code := m.Run()
	os.Exit(code)
}

func TestVersion(t *testing.T) {
	t.Parallel()

	expected, err := os.ReadFile("goldenFiles/version/01.json")
	if err != nil {
		t.Fatal(err)
	}

	tempFilePath := buildTempFilePath(t)
	if err := writeOutput(tempFilePath, hasuraenvBinPath(), "version"); err != nil {
		t.Fatal(err)
	}

	actual, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatal(err)
	}

	jsonEq(t, expected, actual, "time")
}

func TestInit(t *testing.T) {
	t.Cleanup(removeTestConfig)

	if err := exec.Command(hasuraenvBinPath(), "init").Run(); err != nil {
		t.Fatal(err)
	}

	t.Run("expect versions & current are created", func(t *testing.T) {
		expected := []string{"versions", "current"}

		dirEntries, err := os.ReadDir("tmp/test/.hasuraenv")
		if err != nil {
			t.Fatal(err)
		}

		actual := lo.Map(dirEntries, func(entry os.DirEntry, _ int) string {
			return entry.Name()
		})

		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("expect binary is installed", func(t *testing.T) {
		b, err := os.ReadFile("tmp/test/.hasuraenv/versions/default/hasura")

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	t.Run("expect symlink points default", func(t *testing.T) {
		actual, err := os.Readlink("tmp/test/.hasuraenv/current")
		expected := "tmp/test/.hasuraenv/versions/default"

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestLsRemote(t *testing.T) {
	t.Parallel()

	t.Run("expect 30 releases on list when not specifying limit", func(t *testing.T) {
		t.Parallel()

		tempFilePath := buildTempFilePath(t)
		if err := writeOutput(tempFilePath, hasuraenvBinPath(), "ls-remote"); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Fatal(err)
		}

		var output map[string]string
		if err := json.Unmarshal(b, &output); err != nil {
			t.Fatalf("error: %+v data: %s", err, string(b))
		}

		messages := strings.Split(output["msg"], "\n")
		versions := lo.Filter(messages, func(m string, _ int) bool {
			sv := strings.ReplaceAll(m, " ", "")
			sv = strings.Replace(sv, "v", "", 1)
			if _, err := semver.Make(sv); err != nil {
				return false
			}

			return true
		})

		assert.Equal(t, "Latest 30 releases", messages[0])
		assert.Len(t, versions, 30)
	})

	t.Run("expect 10 releases on list when specifying limit as 10", func(t *testing.T) {
		t.Parallel()

		limit := 10
		tempFilePath := buildTempFilePath(t)
		if err := writeOutput(tempFilePath, hasuraenvBinPath(), "ls-remote", "--limit", strconv.Itoa(limit)); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Fatal(err)
		}

		var output map[string]string
		if err := json.Unmarshal(b, &output); err != nil {
			t.Fatalf("error: %+v data: %s", err, string(b))
		}

		messages := strings.Split(output["msg"], "\n")
		versions := lo.Filter(messages, func(m string, _ int) bool {
			sv := strings.ReplaceAll(m, " ", "")
			sv = strings.Replace(sv, "v", "", 1)
			if _, err := semver.Make(sv); err != nil {
				return false
			}

			return true
		})

		assert.Equal(t, "Latest "+strconv.Itoa(limit)+" releases", messages[0])
		assert.Len(t, versions, limit)
	})

	t.Run("expect 110 releases on list when specifying limit as 110", func(t *testing.T) {
		t.Parallel()

		limit := 110
		tempFilePath := buildTempFilePath(t)
		if err := writeOutput(tempFilePath, hasuraenvBinPath(), "ls-remote", "--limit", strconv.Itoa(limit)); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Fatal(err)
		}

		var output map[string]string
		if err := json.Unmarshal(b, &output); err != nil {
			t.Fatalf("error: %+v data: %s", err, string(b))
		}

		messages := strings.Split(output["msg"], "\n")
		versions := lo.Filter(messages, func(m string, _ int) bool {
			sv := strings.ReplaceAll(m, " ", "")
			sv = strings.Replace(sv, "v", "", 1)
			if _, err := semver.Make(sv); err != nil {
				return false
			}

			return true
		})

		assert.Equal(t, "Latest "+strconv.Itoa(limit)+" releases", messages[0])
		assert.Len(t, versions, limit)
	})
}

func TestInstall(t *testing.T) {
	t.Cleanup(removeTestConfig)

	if err := exec.Command(hasuraenvBinPath(), "init").Run(); err != nil {
		t.Fatal(err)
	}

	t.Run("expect v2.13.0 installed", func(t *testing.T) {
		if err := exec.Command(hasuraenvBinPath(), "install", "v2.13.0").Run(); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile("tmp/test/.hasuraenv/versions/v2.13.0/hasura")

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	t.Run("expect an error when specifying non-exist version", func(t *testing.T) {
		err := exec.Command(hasuraenvBinPath(), "install", "v2.0.99").Run()
		if err == nil {
			t.Fatal("v2.0.99 doesn't exist")
		}

		assert.Equal(t, "exit status 1", err.Error())
	})
}

func TestLs(t *testing.T) {
	t.Cleanup(removeTestConfig)

	if err := exec.Command(hasuraenvBinPath(), "init").Run(); err != nil {
		t.Fatal(err)
	}

	if err := exec.Command(hasuraenvBinPath(), "install", "v2.1.0").Run(); err != nil {
		t.Fatal(err)
	}

	if err := exec.Command(hasuraenvBinPath(), "install", "v2.13.0").Run(); err != nil {
		t.Fatal(err)
	}

	tempFilePath := buildTempFilePath(t)
	if err := writeOutput(tempFilePath, hasuraenvBinPath(), "ls"); err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatal(err)
	}

	var output map[string]string
	if err := json.Unmarshal(b, &output); err != nil {
		t.Fatalf("error: %+v data: %s", err, string(b))
	}

	messages := strings.Split(output["msg"], "\n")
	versions := lo.Filter(messages, func(m string, _ int) bool {
		sv := strings.ReplaceAll(m, " ", "")
		sv = strings.Replace(sv, "v", "", 1)
		if _, err := semver.Make(sv); err != nil {
			return false
		}

		return true
	})

	assert.Equal(t, "Installed hasura cli", messages[0])
	assert.Len(t, versions, 2)
}

func TestUse(t *testing.T) {
	t.Cleanup(removeTestConfig)

	if err := exec.Command(hasuraenvBinPath(), "init").Run(); err != nil {
		t.Fatal(err)
	}

	// FIXME: telemetry notice on GitHub Actions
	// NOTE: https://github.com/kmtym1998/hasuraenv/actions/runs/3259113829/jobs/5351709250
	if err := exec.Command(currentHasuraBinPath(), "version", "--skip-update-check").Run(); err != nil {
		t.Fatal(err)
	}

	t.Run("expect switch v2.1.0", func(t *testing.T) {
		expectedHasuraCLIVersion := "v2.1.0"
		if err := exec.Command(hasuraenvBinPath(), "install", expectedHasuraCLIVersion).Run(); err != nil {
			t.Fatal(err)
		}

		if err := exec.Command(hasuraenvBinPath(), "use", expectedHasuraCLIVersion).Run(); err != nil {
			t.Fatal(err)
		}

		tempFilePath := buildTempFilePath(t)
		if err := writeOutput(tempFilePath, currentHasuraBinPath(), "version", "--skip-update-check"); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Fatal(err)
		}

		var output map[string]string
		if err := json.Unmarshal(b, &output); err != nil {
			t.Fatalf("error: %+v data: %s", err, string(b))
		}

		assert.Equal(t, expectedHasuraCLIVersion, output["version"])
	})

	t.Run("expect switch v2.13.0", func(t *testing.T) {
		expectedHasuraCLIVersion := "v2.13.0"
		if err := exec.Command(hasuraenvBinPath(), "install", expectedHasuraCLIVersion).Run(); err != nil {
			t.Fatal(err)
		}

		if err := exec.Command(hasuraenvBinPath(), "use", expectedHasuraCLIVersion).Run(); err != nil {
			t.Fatal(err)
		}

		tempFilePath := buildTempFilePath(t)
		if err := writeOutput(tempFilePath, currentHasuraBinPath(), "version", "--skip-update-check"); err != nil {
			t.Fatal(err)
		}

		b, err := os.ReadFile(tempFilePath)
		if err != nil {
			t.Fatal(err)
		}

		var output map[string]string
		if err := json.Unmarshal(b, &output); err != nil {
			t.Fatalf("error: %+v data: %s", err, string(b))
		}

		assert.Equal(t, expectedHasuraCLIVersion, output["version"])
	})
}
