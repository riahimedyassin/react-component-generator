package config

import (
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/spf13/viper"
)

var (
	GlobalConfig Config // Global app configuration
	defaults     = map[ConfigTokens]any{
		DEFAULT_EXPORT:    false,
		PROJECT_BASE:      enums.TYPESCRIPT,
		PROJECT_STRUCTURE: "feature-based",
		STYLING:           enums.CSS,
		TEST:              false,
		TYPE:              enums.FUNCTIONAL,
	}
)

func setDefaults() {
	for key, value := range defaults {
		viper.SetDefault(string(key), value)
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
