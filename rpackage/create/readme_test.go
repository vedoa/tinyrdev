package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadme(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "readme-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	pkg := PackageMetadata{
		Name: "test",
		Dir:  tmpDir,
	}

	err = readme(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.README)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.README, err)
	}

	expected := readmeContent(pkg.Name)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.README)
	}
}
