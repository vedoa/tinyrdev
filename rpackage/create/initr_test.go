package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitr(t *testing.T) {

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "initr-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	pkg := PackageMetadata{
		Dir: tmpDir,
	}

	// Call the function
	err = initr(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.InitR)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.InitR, err)
	}
	if string(content) != initrContent {
		t.Errorf("%s content mismatch.", files.InitR)
	}
}
