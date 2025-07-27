package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// rprofileContent content of the file
func rprofileContent(pppm string) string {
	return fmt.Sprintf(`options(repos=c("%s"))
source("init.R")
	`,
		pppm)
}

// rprofile creates the .Rprofile file
func rprofile(pppm string, packagePath string) error {
	s := files.Rprofile
	if err := os.WriteFile(filepath.Join(packagePath, s), []byte(rprofileContent(pppm)), 0644); err != nil {
		return err
	}
	return inform(s)
}
