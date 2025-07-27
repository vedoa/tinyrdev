package create

import (
	"fmt"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {

	tmpDir, err := os.MkdirTemp("", "execute-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name        string
		args        []string
		expectError bool
	}{
		{
			name:        "No arguments",
			args:        []string{},
			expectError: false,
		},
		{
			name:        "Version flag only",
			args:        []string{"-v"},
			expectError: false,
		},
		{
			name:        "Missing required name flag",
			args:        []string{"-title", "Test Package"},
			expectError: true,
		},
		{
			name: "All required flags",
			args: []string{
				"-name", "testpkg",
				"-title", "Test Package",
				"-version", "0.1.0",
				"-author", "Tester",
				"-maintainer", "Tester",
				"-email", "test@example.com",
				"-description", "A test package",
				"-license", "MIT",
				"-dir", tmpDir,
				"-createdir=false",
				"-image", "ghcr.io/rocker-org/devcontainer/r-ver:4",
				"-pppm", "https://packagemanager.posit.co/cran/__linux__/noble/latest",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name + " --START--")
			err := Execute("0.1.0", tt.args)
			if (err != nil) != tt.expectError {
				t.Errorf("Execute() error = %v, expectError %v", err, tt.expectError)
			}
			fmt.Println(tt.name + " --END--")
		})
	}
}
