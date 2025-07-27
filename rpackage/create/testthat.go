package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// testthatContent content of the file
func testthatContent(packageName string) string {
	return fmt.Sprintf(`library(testthat)
library(%s)

test_check("%s")
`, packageName, packageName)
}

// testthat holds the information on the testthat.R file
func testthat(packageName string, packagePath string) error {
	s := files.Testthat
	if err := os.WriteFile(filepath.Join(packagePath, "tests", s), []byte(testthatContent(packageName)), 0644); err != nil {
		return err
	}
	return inform(s)
}
