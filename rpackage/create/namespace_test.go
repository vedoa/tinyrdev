package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNamespace(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "namespace-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	pkg := PackageMetadata{
		Dir: tmpDir,
	}

	// Call the function
	err = namespace(pkg)
	if err != nil {
		t.Fatalf("Expected nil found %s", err.Error())
	}

	// Check file exists
	testFile := filepath.Join(pkg.Dir, files.NAMESPACE)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s not found", testFile)
	}

	// Check if content matches
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", files.NAMESPACE, err)
	}
	if string(content) != namespaceContent {
		t.Errorf("%s content mismatch.", files.NAMESPACE)
	}
}
