package create

import (
	"os"
	"path/filepath"
)

// rprojContent content of the file
const rprojContent = `Version: 1.0

RestoreWorkspace: Default
SaveWorkspace: Default
AlwaysSaveHistory: Default

EnableCodeIndexing: Yes
UseSpacesForTab: Yes
NumSpacesForTab: 2
Encoding: UTF-8

RnwWeave: Sweave
LaTeX: pdfLaTeX

AutoAppendNewline: Yes
StripTrailingWhitespace: Yes

BuildType: Package
PackageUseDevtools: Yes
PackageInstallArgs: --no-multiarch --with-keep.source
PackageRoxygenize: rd,collate,namespace
`

// rproj holds the information on the .Rproj file
func rproj(meta PackageMetadata) error {
	s := files.Rproj
	if err := os.WriteFile(filepath.Join(meta.Dir, meta.Name+s), []byte(rprojContent), 0644); err != nil {
		return err
	}
	return inform(s)
}
