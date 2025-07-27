package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// descriptionContent content of the file
func descriptionContent(meta PackageMetadata) string {
	return fmt.Sprintf(`Package: %s
Type: Package
Title: %s
Version: %s
Author: %s
Maintainer: %s <%s>
Description: %s
License: %s
Encoding: UTF-8
LazyData: true
Suggests: 
	testthat,
	roxygen2
Roxygen: list(markdown = TRUE)
RoxygenNote: 7.3.2
`,
		meta.Name,
		meta.Title,
		meta.Version,
		meta.Author,
		meta.Maintainer,
		meta.Email,
		meta.Description,
		meta.License,
	)
}

// description creates the DESCRIPTION file
func description(meta PackageMetadata, packagePath string) error {
	s := files.DESCRIPTION
	if err := os.WriteFile(filepath.Join(packagePath, s), []byte(descriptionContent(meta)), 0644); err != nil {
		return err
	}
	return inform(s)
}
