package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDevcontainer(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "devcontainer-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the .devcontainer directory
	devcontainerDir := filepath.Join(tmpDir, dirs.Devcontainer)
	if err := os.MkdirAll(devcontainerDir, 0755); err != nil {
		t.Fatalf("Failed to create tests dir: %v", err)
	}

	// Fake image name
	img := "ghcr.io/rocker-org/devcontainer/r-ver:4"

	err = devcontainer(img, tmpDir)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(devcontainerDir, files.DevcontainerJSON)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.DevcontainerJSON, err)
	}

	expected := devcontainerContent(img)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.DevcontainerJSON)
	}
}
