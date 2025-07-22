package config

type Config struct {
	core     coreConfig
	linting  lintingConfig
	state    stateConfig
	execPath string
}

func NewConfig(execPath string) *Config {
	return &Config{
		execPath: execPath,
	}
}

func (c *Config) Load() error {
	return nil
}
