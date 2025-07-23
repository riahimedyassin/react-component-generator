package component_generator

import "github.com/riahimedyassin/react-component-generator/config"

type ComponentGenerator struct {
	config *config.GlobalConfig
}

func NewComponentGenerator(config *config.GlobalConfig) *ComponentGenerator {
	return &ComponentGenerator{
		config: config,
	}
}
