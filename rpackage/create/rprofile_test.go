package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRprofile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "rprofile-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	// Call the function
	pkg := PackageMetadata{
		PPPM: "https://packagemanager.posit.co/cran/__linux__/noble/latest",
		Dir:  tmpDir,
	}

	err = rprofile(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.Rprofile)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.Rprofile, err)
	}

	expected := rprofileContent(pkg.PPPM)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.Rprofile)
	}
}
