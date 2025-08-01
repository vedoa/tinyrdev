package create

import (
	"os"
	"path/filepath"
)

// gitignoreContent content of the file
const gitignoreContent = `// # Retrieved 2017-Oct-12 from https://github.com/github/gitignore/blob/master/R.gitignore
// # Licensed under CC0-1.0 https://github.com/github/gitignore/blob/master/LICENSE

# History files
.Rhistory
.Rapp.history

# Session Data files
.RData

# Example code in package build process
*-Ex.R

# Output files from R CMD build
/*.tar.gz

# Output files from R CMD check
/*.Rcheck/

# RStudio files
.Rproj.user/

# produced vignettes
vignettes/*.html
vignettes/*.pdf

# OAuth2 token, see https://github.com/hadley/httr/releases/tag/v0.3
.httr-oauth

# knitr and R markdown default cache directories
/*_cache/
/cache/

# Temporary files created by R markdown
*.utf8.md
*.knit.md

## Retrieved 2023-11-16 from https://github.com/github/gitignore/blob/main/C++.gitignore

# Prerequisites
*.d

# Compiled Object files
*.slo
*.lo
*.o
*.obj

# Precompiled Headers
*.gch
*.pch

# Compiled Dynamic libraries
*.so
*.dylib
*.dll

# Fortran module files
*.mod
*.smod

# Compiled Static libraries
*.lai
*.la
*.a
*.lib

# Executables
*.exe
*.out
*.app

## custom tinyrdev
/env
`

// gitignore creates the .gitignore file
func gitignore(meta PackageMetadata) error {
	s := files.Gitignore
	if err := os.WriteFile(filepath.Join(meta.Dir, s), []byte(gitignoreContent), 0644); err != nil {
		return err
	}
	return inform(s)
}
