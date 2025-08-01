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

	pkg := PackageMetadata{
		Dir:  tmpDir,
		Name: "test",
	}

	err = gitignore(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.Gitignore)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.Gitignore, err)
	}

	if string(content) != gitignoreContent {
		t.Errorf("%s content mismatch.", files.Gitignore)
	}

}
