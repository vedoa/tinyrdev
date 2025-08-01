package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNews(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "news-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	pkg := PackageMetadata{
		Name:    "test",
		Version: "1.0.0",
		Dir:     tmpDir,
	}

	err = news(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.NEWS)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.NEWS, err)
	}

	expected := newsContent(pkg.Name, pkg.Version)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.NEWS)
	}
}
