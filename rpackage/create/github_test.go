package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGithub(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "github-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the .github/workflows directory
	githubDir := filepath.Join(tmpDir, dirs.Github_Workflows)
	if err := os.MkdirAll(githubDir, 0755); err != nil {
		t.Fatalf("Failed to create tests dir: %v", err)
	}

	pkg := PackageMetadata{
		Dir:   tmpDir,
		Image: "ghcr.io/rocker-org/devcontainer/r-ver:4",
	}

	err = github(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(githubDir, files.PackageYML)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.PackageYML, err)
	}

	expected := githubContent(pkg.Image)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.PackageYML)
	}
}
