package e2e

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func hasuraenvBinPath() (p string) {
	p = os.Getenv("HASURAENV_TEST_CLI_PATH")
	if p != "" {
		abs, err := filepath.Abs(p)
		if err != nil {
			panic(err)
		}

		p = abs

		return
	}

	b, err := exec.Command("which", "hasuraenv").Output()
	if err != nil {
		log.Fatal(err)
	}

	p = strings.ReplaceAll(string(b), "\n", "")

	if p != "" {
		return
	}

	log.Fatal("Set 'HASURAENV_TEST_CLI_PATH' or add PATH to hasuraenv executable file for e2e testing.")

	return
}

func buildTempFilePath(t *testing.T) string {
	return fmt.Sprintf("%s/%s.txt", t.TempDir(), uuid.New().String())
}

func execSubCommand(tempFilePath string, arg ...string) error {
	command := fmt.Sprintf("%s %s > %s",
		hasuraenvBinPath(), strings.Join(arg, " "), tempFilePath,
	)

	println(command)

	return exec.Command("sh", "-c", command).Run()
}

// NOTE: modification of stretchr/testify/assert.JSONEq
//       https://github.com/stretchr/testify/blob/181cea6eab8b2de7071383eca4be32a424db38dd/assert/assertions.go#L1607-L1622

type tHelper interface {
	Helper()
}

func jsonEq(t assert.TestingT, expected, actual []byte, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	var expectedJSONAsMap, actualJSONAsMap map[string]any

	if err := json.Unmarshal([]byte(expected), &expectedJSONAsMap); err != nil {
		return assert.Fail(t, fmt.Sprintf("Expected value ('%s') is not valid json.\nJSON parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	if err := json.Unmarshal([]byte(actual), &actualJSONAsMap); err != nil {
		return assert.Fail(t, fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	delete(expectedJSONAsMap, "time")
	delete(actualJSONAsMap, "time")

	expectedJSONAsBytes, err := json.Marshal(expectedJSONAsMap)
	if err != nil {
		assert.Fail(t, "Marshal error: %+v", expectedJSONAsMap)
	}

	actualJSONAsBytes, err := json.Marshal(actualJSONAsMap)
	if err != nil {
		assert.Fail(t, "Marshal error: %+v", actualJSONAsMap)
	}

	var expectedJSONAsAny, actualJSONAsAny any

	json.Unmarshal(expectedJSONAsBytes, &expectedJSONAsAny)
	json.Unmarshal(actualJSONAsBytes, &actualJSONAsAny)

	return assert.Equal(t, expectedJSONAsAny, actualJSONAsAny, msgAndArgs...)
}
