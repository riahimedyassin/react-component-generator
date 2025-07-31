package page_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
)

// The page and the component do not have any difference except for the route generation process. The seperation in this phase isn't justified, but in future realeases this could be helpfull to manage pages and components in seperate ways.
type PageWrapper struct {
	compWrapper *component_generator.ComponentWrapper
	config      *config.Config
	name, path  string
	flags       *FlagsOptions
}

func NewPageWrapper(
	compWrapper *component_generator.ComponentWrapper,
	name, path string,
	config *config.Config,
	flags *FlagsOptions,
) *PageWrapper {
	return &PageWrapper{
		compWrapper: compWrapper,
		config:      config,
		name:        name,
		path:        path,
		flags:       flags,
	}
}

func (w *PageWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	return w.compWrapper.GetDefiners()
}

func (h *PageWrapper) GetEditors() []generator_interfaces.Editor {
	editors := []generator_interfaces.Editor{
		NewAppEditor(h.name, h.path, h.config, h.flags),
		NewRouteEditor(h.name, h.path, h.config, h.flags),
	}
	return editors
}
