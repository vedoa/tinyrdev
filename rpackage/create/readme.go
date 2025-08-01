package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// newsContent content of the file
func readmeContent(packageName string) string {
	return fmt.Sprintf(`# %s

A brief description of what your R package does and its purpose.

## Getting Started

### Prerequisites

TODO

### Installation

TODO

## Usage

TODO
  
## Contributing

TODO
`,
		packageName)
}

// news creates the NEWS.md file
func readme(meta PackageMetadata) error {
	s := files.README
	if err := os.WriteFile(filepath.Join(meta.Dir, s), []byte(readmeContent(meta.Name)), 0644); err != nil {
		return err
	}
	return inform(s)
}
