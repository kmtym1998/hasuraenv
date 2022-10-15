package e2e

import (
	"os"
	"os/exec"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
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
	if err := writeOutput(tempFilePath, "version"); err != nil {
		t.Fatal(err)
	}

	actual, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatal(err)
	}

	jsonEq(t, expected, actual, "time")
}

func TestInit(t *testing.T) {
	t.Parallel()

	if err := exec.Command(hasuraenvBinPath(), "init").Run(); err != nil {
		t.Fatal(err)
	}

	t.Run("expect versions & current are created", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
		b, err := os.ReadFile("tmp/test/.hasuraenv/versions/default/hasura")

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})

	t.Run("expect symlink points default", func(t *testing.T) {
		t.Parallel()
		actual, err := os.Readlink("tmp/test/.hasuraenv/current")
		expected := "tmp/test/.hasuraenv/versions/default"

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
