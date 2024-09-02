package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path, name string) (string, error) {
	file, err := os.ReadFile(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func WriteFile(path, name, content string) error {
	return os.WriteFile(fmt.Sprintf("%s/%s", path, name), []byte(content), os.ModeAppend)
}

func ParseJson(content string, target interface{}) error {
	return json.Unmarshal([]byte(content), target)
}
