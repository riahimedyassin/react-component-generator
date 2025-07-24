package files

import (
	"fmt"
	"io/fs"
	"os"
)

func WriteFile(path, name, extension, content string) error {
	fullPath := fmt.Sprintf("%s\\%s.%s", path, name, extension)
	fmt.Print(fullPath)
	err := os.WriteFile(fullPath, []byte(content), fs.ModePerm)
	if err != nil {
		fmt.Print("write file error")
	}
	return err
}

func ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}
