package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDescription(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "description-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Fake metadata
	pkg := PackageMetadata{
		Name:        "mypackage",
		Title:       "An Example R Package",
		Version:     "0.1.0",
		Author:      "Jane Doe",
		Maintainer:  "Jane Doe",
		Email:       "jane@example.com",
		Description: "This package demonstrates metadata usage.",
		License:     "MIT",
		Image:       "ghcr.io/rocker-org/devcontainer/r-ver:4",
		CreateDir:   true,
		PPPM:        "https://packagemanager.posit.co/cran/__linux__/noble/latest",
	}

	err = description(pkg, tmpDir)
	if err != nil {
		t.Fatalf("Expected nill found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(tmpDir, files.DESCRIPTION)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.DESCRIPTION, err)
	}

	expected := descriptionContent(pkg)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.DESCRIPTION)
	}
}
