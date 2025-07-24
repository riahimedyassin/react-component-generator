package files

import (
	"fmt"
	"os"
)

func WriteFile(path, name, extension, content string) error {
	return os.WriteFile(fmt.Sprintf("%s\\%s", path, name), []byte(content), os.ModeAppend)
}

func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func appendFile(path, content string) error   // APPEND Flag
func overrideFile(path, content string) error // OVERRIDE Flag
func createFile(path, content string) error   // CREATE Flag
