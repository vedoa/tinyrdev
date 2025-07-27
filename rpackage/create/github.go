package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// githubContent content of the file
func githubContent(image string) string {
	return fmt.Sprintf(`name: R Package

on:
  push:
    branches: ["*"]
  pull_request:
    branches: ["*"]

jobs:
  r-package:
    runs-on: ubuntu-latest

    container:
      image: %s

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Install dependencies of the package
        run: |
          Rscript -e 'devtools::install_deps(dependencies = TRUE)' 

      - name: Document the package
        run: |
          Rscript -e 'devtools::document()'

      - name: Check the package
        run: |
          Rscript -e 'devtools::check()'

`,
		image)
}

// github holds the information on the package.yml file
func github(image string, packagePath string) error {
	s := files.PackageYML
	if err := os.WriteFile(filepath.Join(packagePath, dirs.Github_Workflows, s), []byte(githubContent(image)), 0644); err != nil {
		return err
	}
	return inform(s)
}
