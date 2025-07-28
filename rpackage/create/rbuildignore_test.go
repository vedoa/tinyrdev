package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRbuildignore(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "rbuildignore-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	// Call the function
	err = rbuildignore(tmpDir)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(tmpDir, files.Rbuildignore)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.Rbuildignore, err)
	}
	if string(content) != rbuildignoreContent {
		t.Errorf("%s content mismatch.", files.Rbuildignore)
	}
}
