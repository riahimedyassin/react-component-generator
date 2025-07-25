package config

import "github.com/riahimedyassin/react-component-generator/config/enums"

type componentConfig struct {
	Type          enums.ComponentType
	DefaultExport bool
	Test          bool
	WithStyling   bool
	WithTest      bool
}

type projectConfig struct {
	Styling   enums.Styling
	Base      enums.ProjectBase
	Structure string
}

type Config struct {
	Component componentConfig
	Project   projectConfig
}
