package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

type TestGenerator struct {
	name    string
	path    string
	config  *config.Config
	options *FlagsOptions
}

func NewTestGenerator(name, path string, config *config.Config, options *FlagsOptions) *TestGenerator {
	return &TestGenerator{
		name:    name,
		path:    path,
		config:  config,
		options: options,
	}
}

// todo : spec file content
func (g *TestGenerator) Generate() error {
	if g.config.Component.WithTest {
		if err := files.WriteFile(g.path, g.name, g.getExtension(), ""); err != nil {
			return err
		}
	}
	return nil
}

func (g *TestGenerator) getExtension() string {
	if g.config.Project.Base == enums.TYPESCRIPT {
		return "spec.ts"
	}
	return "spec.js"
}
