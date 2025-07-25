package lib

import (
	"path/filepath"
)

func GetTemplatePath(template string) string {
	// exePath, _ := os.Executable()
	// exeDir := filepath.Dir(exePath)
	return filepath.Join("./templates", template)
}
