package create

import (
	"errors"
	"flag"
	"fmt"
)

// Execute contains the current main flags. Would be the create subcommand if the cli grows
func Execute(v string, args []string) error {

	fs := flag.NewFlagSet("create", flag.ContinueOnError)

	// Define flags
	name := fs.String("name", "", "Name of the R package (required)")
	title := fs.String("title", "What the Package Does (One Line, Title Case)", "Title of the package")
	version := fs.String("version", "0.1.0", "Version of the package")
	author := fs.String("author", "Your Name", "Author of the package")
	maintainer := fs.String("maintainer", "Your Name", "Maintainer of the package")
	email := fs.String("email", "your@email.com", "Maintainer email")
	description := fs.String("description", "A short description of the package.", "Description of the package")
	license := fs.String("license", "MIT", "License type")
	dir := fs.String("dir", ".", "Target directory to create the package in")
	createdir := fs.Bool("createdir", true, "Should a folder with the package name be created")
	image := fs.String("image", "ghcr.io/rocker-org/devcontainer/r-ver:4", "Devcontainer image")
	pppm := fs.String("pppm", "https://packagemanager.posit.co/cran/__linux__/noble/latest", "Which Posit Public Package Manager date should be used?")
	vFlag := fs.Bool("v", false, "Print version")

	// Show help if no arguments are provided
	if len(args) == 0 {
		fmt.Println("Version:", v)
		fs.Usage()
		return nil
	}

	fs.Parse(args)

	// If only version is requested print and exit
	if len(args) == 1 && *vFlag {
		fmt.Println(v)
		return nil
	}

	// Validate required fields
	if *name == "" {
		return errors.New("package name is required. Use -name to specify it")
	}

	// Create metadata struct
	meta := PackageMetadata{
		Name:        *name,
		Title:       *title,
		Version:     *version,
		Author:      *author,
		Maintainer:  *maintainer,
		Email:       *email,
		Description: *description,
		License:     *license,
		Image:       *image,
		CreateDir:   *createdir,
		PPPM:        *pppm,
	}

	// Create the package
	err := CreatePackage(meta, *dir)
	if err != nil {
		return errors.New("Failed to create package: " + err.Error())
	}
	fmt.Printf("R package '%s' created successfully at %s\n", *name, *dir)
	return nil
}
