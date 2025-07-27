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
	pppm := "https://packagemanager.posit.co/cran/__linux__/noble/latest"

	err = rprofile(pppm, tmpDir)
	if err != nil {
		t.Fatalf("Expected nill found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(tmpDir, files.Rprofile)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.Rprofile, err)
	}

	expected := rprofileContent(pppm)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.Rprofile)
	}
}
