package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type ComponentWrapper struct {
	name, path string
	config     *config.Config
	flags      *FlagsOptions
}

func NewComponentWrapper(
	name, path string,
	config *config.Config,
	flags *FlagsOptions,
) *ComponentWrapper {
	return &ComponentWrapper{
		name:   name,
		path:   path,
		config: config,
		flags:  flags,
	}
}

func (c *ComponentWrapper) GetDefiners() []generator.FileSpecDefiner {
	definers := []generator.FileSpecDefiner{
		newComponentFileSpecGenerator(c.name, c.path, c.config, c.flags),
	}
	if c.config.Component.WithStyling {
		definers = append(definers, newStyleGenerator(c.name, c.path, c.config, c.flags))
	}
	if c.config.Component.WithTest {
		definers = append(definers, newTestGenerator(c.name, c.path, c.config, c.flags))
	}
	return definers
}
