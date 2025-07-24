package config

import (
	"github.com/spf13/viper"
)

var (
	config   Config
	defaults = map[string]any{
		"core.structure": "",
		"core.template":  "tsx",
		"core.styling":   "tailwind",
		"state.type":     "zustand",
		"linting.eslint": true,
		"core.prettier":  true,
	}
)

func setDefaults() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

func Load(cwd string) (*Config, error) {
	setDefaults()
	viper.AddConfigPath(cwd)
	viper.SetConfigName("rgconfig")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
