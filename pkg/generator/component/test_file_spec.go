package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type testGenerator struct {
	name   string
	path   string
	config *config.Config
	flags  *FlagsOptions
}

func newTestGenerator(execContent *generator_models.ExecContenxt[FlagsOptions]) *testGenerator {
	return &testGenerator{
		name:   execContent.Filename,
		path:   execContent.Filepath,
		config: execContent.Config,
		flags:  &execContent.Flags,
	}
}

// todo :  generate testing content
func (g *testGenerator) GetFileSpec() (*generator_models.FileSpec, error) {
	if g.config.Component.WithTest {
		return generator_models.NewFileSpec(g.name, g.path, g.getExtension(), ""), nil
	}
	return nil, rg_errors.NewIgnoreDefinerError()
}

func (g *testGenerator) getExtension() string {
	if g.config.Project.Base == enums.TYPESCRIPT {
		return "spec.ts"
	}
	return "spec.js"
}
