package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRproj(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "rproj-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	pkg := PackageMetadata{
		Name: "mypackage",
		Dir:  tmpDir,
	}

	err = rproj(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, pkg.Name+files.Rproj)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.Rproj, err)
	}
	if string(content) != rprojContent {
		t.Errorf("%s content mismatch.", files.Rproj)
	}
}
