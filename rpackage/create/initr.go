package create

import (
	"os"
	"path/filepath"
)

// initrContent content of the file
const initrContent = `# set current path
packagePath <- file.path(getwd(), "env")

# check existence
if(!dir.exists(packagePath)){
  # create if necessary
  dir.create(packagePath)
  # inform user
  cat(paste0("Path ", packagePath, " created.\n"))
}

# add to libpaths
.libPaths(packagePath)

# delete from environment
rm(packagePath)

# inform user
cat(paste0("Environments loaded\n", paste0(.libPaths(), collapse = "\n")))
cat(paste0("\nCRAN set to:\n"), paste0(options()$repos, collapse = "\n"))
cat("\n")
`

// initr holds the information on the init.R file
func initr(meta PackageMetadata) error {
	s := files.InitR
	if err := os.WriteFile(filepath.Join(meta.Dir, s), []byte(initrContent), 0644); err != nil {
		return err
	}
	return inform(s)
}
