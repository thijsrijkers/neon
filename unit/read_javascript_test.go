package unit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"neon/runtime"
)

func TestRunScriptsFromFolder(t *testing.T) {
	rt, err := runtime.NewRuntime()
	if err != nil {
		t.Fatalf("Failed to create runtime: %v", err)
	}

	scriptDir := "./scripts"
	entries, err := os.ReadDir(scriptDir)
	if err != nil {
		t.Fatalf("Failed to read scripts folder: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".js") {
			continue
		}

		path := filepath.Join(scriptDir, entry.Name())
		t.Run(entry.Name(), func(t *testing.T) {
			result, err := rt.RunFile(path)
			if err != nil {
				t.Errorf("Script %s failed to run: %v", path, err)
			} else {
				t.Logf("Script %s ran successfully, result: %s", path, result)
			}
		})
	}
}
