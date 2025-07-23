package config

import (
	"encoding/json"

	"github.com/riahimedyassin/react-component-generator/lib/files"
)

type Config struct {
	globalConfig GlobalConfig
	execPath     string
}

func NewConfig(execPath string) *Config {
	return &Config{
		execPath: execPath,
	}
}

func (c *Config) Load() error {
	content, err := files.ReadFile(c.execPath)
	if err != nil {
		return err
	}
	var config GlobalConfig
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		return err
	}
	c.globalConfig = config
	return nil
}

func (c *Config) GenerateConfig(flags string) error {
	return nil
}
