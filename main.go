package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"neon/runtime"
)

func main() {
	rt, err := runtime.NewRuntime()
	if err != nil {
		log.Fatal("Failed to create JS runtime:", err)
	}

	scriptDir := "./scripts"
	entries, err := os.ReadDir(scriptDir)
	if err != nil {
		log.Fatalf("Failed to read scripts folder: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".js") {
			continue
		}

		path := filepath.Join(scriptDir, entry.Name())
		fmt.Printf("▶️ Running %s...\n", path)

		result, err := rt.RunFile(path)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		fmt.Printf("✅ Result: %s\n\n", result)
	}
}
