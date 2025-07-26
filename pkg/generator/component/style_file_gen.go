package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

// genereate styling files
type StyleGenerator struct {
	name    string
	path    string
	config  *config.Config
	options *FlagsOptions
}

func NewStyleGenerator(name, path string, config *config.Config, options *FlagsOptions) *StyleGenerator {
	return &StyleGenerator{
		name:    name,
		path:    path,
		config:  config,
		options: options,
	}
}

func (g *StyleGenerator) Generate() error {
	switch g.config.Project.Styling {
	case enums.NONE, enums.TAILWIND:
		return nil
	}
	if g.config.Component.WithStyling {
		extension := g.getStyleExtension()
		if err := files.WriteFile(g.path, g.name, extension, ""); err != nil {
			return err
		}
	}
	return nil
}

func (g *StyleGenerator) getStyleExtension() string {
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
