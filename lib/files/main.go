package files

import "os"

func WriteFile(path, content string, flag Flags) error {
	return nil
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
