package config

import "github.com/riahimedyassin/react-component-generator/config/enums"

type componentConfig struct {
	Styling       enums.Styling
	Type          enums.ComponentType
	DefaultExport bool
	Test          bool
}

type projectConfig struct {
	Base      enums.ProjectBase
	Structure string
}

type Config struct {
	Component componentConfig
	Project   projectConfig
}
