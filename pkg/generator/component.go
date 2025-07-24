package generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
)

type ComponentGenerator struct {
	config *config.Config
}

func NewComponentGenerator(config *config.Config) *ComponentGenerator {
	return &ComponentGenerator{
		config: config,
	}
}

func (g *ComponentGenerator) Generate(cwd string, name string) error {
	return nil
}

func (g *ComponentGenerator) generateComponent(cwd, name string, template enums.Template) error {
	return nil
}

func (g *ComponentGenerator) getTemplate(name string, template enums.Template) ([]byte, error) {
	return nil, nil
}
