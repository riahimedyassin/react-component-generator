package hooks_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type hookTestGenerator struct {
	name    string
	path    string
	config  *config.Config
	options *FlagsOptions
}

func newHookTestGenerator(name, path string, config *config.Config, options *FlagsOptions) *hookTestGenerator {
	return &hookTestGenerator{
		name:    name,
		path:    path,
		config:  config,
		options: options,
	}
}

// todo : spec file content
func (g *hookTestGenerator) GetFileSpec() (*generator.FileSpec, error) {
	if g.config.Component.WithTest {
		return generator.NewFileSpec(g.name, g.path, g.getExtension(), ""), nil
	}
	return nil, rg_errors.NewIgnoreDefinerError()
}

func (g *hookTestGenerator) getExtension() string {
	if g.config.Project.Base == enums.TYPESCRIPT {
		return "spec.ts"
	}
	return "spec.js"
}
