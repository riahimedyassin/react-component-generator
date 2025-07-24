package config

import "github.com/riahimedyassin/react-component-generator/config/enums"

type coreConfig struct {
	Structure string         `json:"structure"`
	Template  enums.Template `json:"template"`
	Styling   string         `json:"styling"`
}

type stateConfig struct {
	Type string `json:"type"`
}

type lintingConfig struct {
	Eslint   bool `json:"eslint"`
	Prettier bool `json:"prettier"`
}

type componentConfig struct {
	WithStyling bool
	WithTest    bool
	Type        enums.ComponentType
}

type Config struct {
	Core      coreConfig      `json:"core"`
	State     stateConfig     `json:"state"`
	Linting   lintingConfig   `json:"linting"`
	Component componentConfig `json:"component"`
}
