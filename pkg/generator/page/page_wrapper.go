package page_generator

import (
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
)

// The page and the component do not have any difference except for the route generation process. The seperation in this phase isn't justified, but in future realeases this could be helpfull to manage pages and components in seperate ways.
type PageWrapper struct {
	compWrapper *component_generator.ComponentWrapper
}

func NewPageWrapper(compWrapper *component_generator.ComponentWrapper) *PageWrapper {
	return &PageWrapper{
		compWrapper: compWrapper,
	}
}

func (w *PageWrapper) GetDefiners() []generator.FileSpecDefiner {
	return w.compWrapper.GetDefiners()
}
