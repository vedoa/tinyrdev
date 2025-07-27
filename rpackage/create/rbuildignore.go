package create

import (
	"os"
	"path/filepath"
)

// rbuildignore content of the file
const rbuildignoreContent = `^pkgdown$
^CRAN-RELEASE$
visual_test
^.*\.Rproj$
^\.Rproj\.user$
^/\.gitattributes$
^inst/web$
^notes$
^cran-comments.md$
^data-raw$
^CONTRIBUTING\.md$
^GOVERNANCE\.md$
^ISSUE_TEMPLATE\.md$
^CODE_OF_CONDUCT\.md$
^NEWS$
^NEWS.md$
^revdep$
^\.httr-oauth$
^codecov\.yml$
^_pkgdown\.yml$
^icons$
^README\.Rmd$
^README-.*\.png$
^logo\.png$
^\.github$
^vignettes/profilings
^vignettes/profiling.Rmd$
^cran-comments\.md$
^LICENSE\.md$
^vignettes/articles$
^CRAN-SUBMISSION$
^docs$
^\.devcontainer$
`

// rbuildignore holds the information on the .Rbuildignore file
func rbuildignore(packagePath string) error {
	s := files.Rbuildignore
	if err := os.WriteFile(filepath.Join(packagePath, s), []byte(rbuildignoreContent), 0644); err != nil {
		return err
	}
	return inform(s)
}
