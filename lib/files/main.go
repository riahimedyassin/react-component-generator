package files

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileSystem struct {
}

var fileSystem *FileSystem

func NewFileSystem() *FileSystem {
	if fileSystem == nil {
		fileSystem = &FileSystem{}
	}
	return fileSystem
}

func (f *FileSystem) WriteFile(path, name, extension, content string) error {
	fullPath := filepath.Join(path, name+"."+extension)
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}
	err := os.WriteFile(fullPath, []byte(content), fs.ModePerm)
	if err != nil {
		fmt.Print("write file error")
	}
	return err
}

func (f *FileSystem) ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (f *FileSystem) DeleteFile(path, name, extension string) error {
	fullFilePath := fmt.Sprintf("%s/%s.%s", path, name, extension)
	return os.Remove(fullFilePath)
}
