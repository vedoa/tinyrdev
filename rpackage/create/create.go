package create

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// CreatePackage creates the necessary folder structure
func CreatePackage(meta PackageMetadata, dir string) error {
	packagePath := dir
	// main package path
	if meta.CreateDir {
		packagePath = filepath.Join(packagePath, meta.Name)
	}
	// Create root directory
	if err := os.MkdirAll(packagePath, 0755); err != nil {
		return fmt.Errorf("creating package directory: %w", err)
	}

	// Create subdirectories
	for _, sub := range dirs.DirNames {
		if err := os.MkdirAll(filepath.Join(packagePath, sub), 0755); err != nil {
			log.Fatal("creating subdirectory " + sub + ":" + err.Error())
		}
		log.Print("creating subdirectory " + sub + "\n")
	}

	//create files
	err := description(meta, packagePath)
	if err != nil {
		return err
	}
	err = namespace(packagePath)
	if err != nil {
		return err
	}
	err = rproj(meta.Name, packagePath)
	if err != nil {
		return err
	}
	err = testthat(meta.Name, packagePath)
	if err != nil {
		return err
	}
	err = gitignore(meta.Name, packagePath)
	if err != nil {
		return err
	}
	err = rbuildignore(packagePath)
	if err != nil {
		return err
	}
	err = devcontainer(meta.Image, packagePath)
	if err != nil {
		return err
	}
	err = initr(packagePath)
	if err != nil {
		return err
	}
	err = rprofile(meta.PPPM, packagePath)
	if err != nil {
		return err
	}
	err = github(meta.Image, packagePath)
	if err != nil {
		return err
	}
	err = news(meta.Name, meta.Version, packagePath)
	if err != nil {
		return err
	}
	err = readme(meta.Name, packagePath)
	if err != nil {
		return err
	}
	return nil
}
