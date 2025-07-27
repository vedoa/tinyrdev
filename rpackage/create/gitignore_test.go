package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGitignore(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "gitignore-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Fake image name
	packageName := "test"

	err = gitignore(packageName, tmpDir)
	if err != nil {
		t.Fatalf("Expected nill found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(tmpDir, files.Gitignore)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.Gitignore, err)
	}

	expected := gitignoreContent(packageName)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.Gitignore)
	}
}
