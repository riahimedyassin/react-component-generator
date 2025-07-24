package config

import (
	"github.com/riahimedyassin/react-component-generator/config/tokens"
	"github.com/spf13/viper"
)

var (
	config   Config
	defaults = map[string]any{
		tokens.STRUCTURE:  "",
		tokens.TEMPLATE:   "tsx",
		tokens.STYLING:    "tailwind",
		tokens.STATE_TYPE: "zustand",
		tokens.ESLINT:     true,
		tokens.PRETTIER:   true,
	}
)

func setDefaults() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

// Called after
func Load(cwd string) error {
	setDefaults()
	viper.AddConfigPath(cwd)
	viper.SetConfigName("rg.config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				return err
			}
		}
		return err
	}
	return viper.Unmarshal(&config)
}
