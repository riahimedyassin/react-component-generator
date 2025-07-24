package generator

import (
	"fmt"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
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
	return g.generateComponent(cwd, name)
}

func (g *ComponentGenerator) generateComponent(cwd, name string) error {
	template, err := g.getTemplate(g.config.Component.Type)
	if err != nil {
		return err
	}
	return files.WriteFile(cwd, name, string(g.config.Core.Template), string(template))
}

func (g *ComponentGenerator) getTemplate(compType enums.ComponentType) ([]byte, error) {
	content, err := files.ReadFile(config.COMP_PATH[compType])
	if err != nil {
		fmt.Print("template not found")
	}
	return content, err
}
