package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
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

func (c *ComponentWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	definers := []generator_interfaces.FileSpecDefiner{
		newComponentFileSpecGenerator(c.name, c.path, c.config, c.flags, files.NewFileSystem()),
	}
	if c.config.Component.WithStyling {
		definers = append(definers, newStyleGenerator(c.name, c.path, c.config, c.flags))
	}
	if c.config.Component.WithTest {
		definers = append(definers, newTestGenerator(c.name, c.path, c.config, c.flags))
	}
	return definers
}

func (h *ComponentWrapper) GetEditors() []generator_interfaces.Editor {
	return nil
}
