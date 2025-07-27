package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// newsContent content of the file
func newsContent(packageName string, packageVersion string) string {
	return fmt.Sprintf(`# %s
## %s %s
	`,
		packageName,
		packageName,
		packageVersion)
}

// news creates the NEWS.md file
func news(packageName string, packageVersion string, packagePath string) error {
	s := files.NEWS
	if err := os.WriteFile(filepath.Join(packagePath, s), []byte(newsContent(packageName, packageVersion)), 0644); err != nil {
		return err
	}
	return inform(s)
}
