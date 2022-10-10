package e2e

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test1(t *testing.T) {
	expected, err := os.ReadFile("goldenFiles/version/01.json")
	if err != nil {
		t.Fatal(err)
	}

	tempFilePath := buildTempFilePath(t)
	if err := execSubCommand(tempFilePath, "version"); err != nil {
		t.Fatal(err)
	}

	actual, err := os.ReadFile(tempFilePath)
	if err != nil {
		t.Fatal(err)
	}

	jsonEq(t, expected, actual)
}
