package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreatePackage creates the necessary folder structure
func CreatePackage(meta PackageMetadata) error {
	// main package path
	if meta.CreateDir {
		meta.Dir = filepath.Join(meta.Dir, meta.Name)
	}
	// Create root directory
	if err := os.MkdirAll(meta.Dir, 0755); err != nil {
		return fmt.Errorf("creating package directory: %w", err)
	}

	// Create subdirectories
	for _, sub := range dirs.DirNames {
		if err := os.MkdirAll(filepath.Join(meta.Dir, sub), 0755); err != nil {
			log.Fatal("creating subdirectory " + sub + ":" + err.Error())
		}
		log.Print("creating subdirectory " + sub + "\n")
	}

	//create files
	tasks := []func(PackageMetadata) error{
		description,
		devcontainer,
		github,
		gitignore,
		initr,
		namespace,
		news,
		rbuildignore,
		readme,
		rprofile,
		rproj,
		testthat,
	}

	for _, task := range tasks {
		if err := task(meta); err != nil {
			return err
		}
	}

	return nil
}
