package lib

import (
	"os"
	"path/filepath"
)

func GetTemplatePath(template string) string {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "templates", template)
}
