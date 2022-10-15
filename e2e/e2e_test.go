package e2e

import (
	"os"
	"testing"
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
