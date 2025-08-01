package create

// PackageMetadata holds the metadata of the package
type PackageMetadata struct {
	Name        string
	Title       string
	Version     string
	Author      string
	Maintainer  string
	Email       string
	Description string
	License     string
	Image       string
	CreateDir   bool
	PPPM        string
	Dir         string
}

// Define the struct
type PackageFiles struct {
	DESCRIPTION      string
	DevcontainerJSON string
	PackageYML       string
	Gitignore        string
	InitR            string
	NAMESPACE        string
	Rbuildignore     string
	Rprofile         string
	Rproj            string
	Testthat         string
	NEWS             string
	README           string
}

// Create a "constant" instance
var files = PackageFiles{
	DESCRIPTION:      "DESCRIPTION",
	DevcontainerJSON: "devcontainer.json",
	PackageYML:       "package.yml",
	Gitignore:        ".gitignore",
	InitR:            "init.R",
	NAMESPACE:        "NAMESPACE",
	Rbuildignore:     ".Rbuildignore",
	Rprofile:         ".Rprofile",
	Rproj:            ".Rproj",
	Testthat:         "testthat.R",
	NEWS:             "NEWS.md",
	README:           "README.md",
}

type PackageDirs struct {
	DirNames         []string
	R                string
	MAN              string
	Tests_Testthat   string
	Tests            string
	Devcontainer     string
	Github_Workflows string
}

// Create a "constant" instance
var dirs = PackageDirs{
	DirNames:         []string{"R", "man", "tests/testthat", ".devcontainer", ".github/workflows"},
	R:                "R",
	MAN:              "man",
	Tests_Testthat:   "tests/testthat",
	Tests:            "tests",
	Devcontainer:     ".devcontainer",
	Github_Workflows: ".github/workflows",
}
