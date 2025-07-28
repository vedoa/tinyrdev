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

	// Call the function
	packageName := "test"
	packageVersion := "1.0.0"
	err = news(packageName, packageVersion, tmpDir)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(tmpDir, files.NEWS)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.NEWS, err)
	}

	expected := newsContent(packageName, packageVersion)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.NEWS)
	}
}
