package create

import (
	"fmt"
	"os"
	"path/filepath"
)

// devcontainerContent content of the file
func devcontainerContent(image string) string {
	return fmt.Sprintf(`{
	"image": "%s",
	"customizations": {
		"vscode": {
			"extensions": [
				"REditorSupport.r", //R
				"streetsidesoftware.code-spell-checker" //code spell checking
			]
		}
	}
}
`,
		image)
}

// devcontainer creates the devcontainer.json file
func devcontainer(image string, packagePath string) error {
	s := files.DevcontainerJSON
	if err := os.WriteFile(filepath.Join(packagePath, dirs.Devcontainer, s), []byte(devcontainerContent(image)), 0644); err != nil {
		return err
	}
	return inform(s)
}
