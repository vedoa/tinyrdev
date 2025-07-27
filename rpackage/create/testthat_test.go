package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTestthat(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "testthat-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the tests directory
	testsDir := filepath.Join(tmpDir, dirs.Tests)
	if err := os.MkdirAll(testsDir, 0755); err != nil {
		t.Fatalf("Failed to create tests dir: %v", err)
	}

	packageName := "mypackage"

	err = testthat(packageName, tmpDir)
	if err != nil {
		t.Fatalf("Expected nill found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(testsDir, files.Testthat)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check content
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", files.Testthat, err)
	}

	expected := testthatContent(packageName)
	if string(content) != expected {
		t.Errorf("%s content mismatch.", files.Testthat)
	}
}
