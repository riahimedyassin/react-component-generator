package config

import (
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/config/tokens"
	"github.com/spf13/viper"
)

var (
	GlobalConfig Config // Global app configuration
	defaults     = map[string]any{
		tokens.STRUCTURE:  "feature-based",
		tokens.TEMPLATE:   "tsx",
		tokens.STYLING:    "tailwind",
		tokens.STATE_TYPE: "zustand",
		tokens.ESLINT:     true,
		tokens.PRETTIER:   true,
		tokens.TYPE:       enums.FUNCTIONNAL,
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
	return viper.Unmarshal(&GlobalConfig)
}
