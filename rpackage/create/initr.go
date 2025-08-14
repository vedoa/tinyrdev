package create

import (
	"os"
	"path/filepath"
)

// initrContent content of the file
const initrContent = `# Get OS
sysName <- Sys.info()[["sysname"]]

if (sysName == "Windows") {
  cat("Running on Windows\n")
} else if (sysName == "Linux") {
  cat("Running on Linux\n")
} else {
  stop(paste0(sysName," not supported"))
}
sysName <- tolower(sysName)

# set current path
packagePath <- file.path(getwd(), "env", sysName)

# check existence
if(!dir.exists(packagePath)){
  # create if necessary
  dir.create(packagePath, recursive = TRUE)
  # inform user
  cat(paste0("Path ", packagePath, " created.\n"))
}

# add to libPaths
.libPaths(packagePath)

# delete from environment
rm(packagePath, sysName)

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
