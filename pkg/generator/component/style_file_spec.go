package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

// genereate styling files
type styleGenerator struct {
	name    string
	path    string
	config  *config.Config
	options *FlagsOptions
}

func newStyleGenerator(name, path string, config *config.Config, options *FlagsOptions) *styleGenerator {
	return &styleGenerator{
		name:    name,
		path:    path,
		config:  config,
		options: options,
	}
}

func (g *styleGenerator) GetFileSpec() (*generator_models.FileSpec, error) {
	switch g.config.Project.Styling {
	case enums.NONE, enums.TAILWIND:
		return nil, rg_errors.NewIgnoreDefinerError()
	}
	if g.config.Component.WithStyling {
		return generator_models.NewFileSpec(g.name, g.path, g.getStyleExtension(), ""), nil
	}
	return nil, rg_errors.NewIgnoreDefinerError()
}

func (g *styleGenerator) getStyleExtension() string {
	extensions := map[enums.Styling]string{
		enums.SCSS: "scss",
		enums.CSS:  "css",
	}
	extension, ok := extensions[g.config.Project.Styling]
	if !ok {
		extension = "css" // default style
	}
	return extension
}
