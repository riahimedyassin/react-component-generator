package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
)

type ComponentGenerator struct {
	config *config.Config
}

func NewComponentGenerator(config *config.Config) *ComponentGenerator {
	return &ComponentGenerator{
		config: config,
	}
}

// todo :  Check if the name is a path and dive/create the target directory.
func (g *ComponentGenerator) Generate(cwd string, name string) error {
	return nil
}
