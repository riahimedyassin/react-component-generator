package page_generator

import (
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

// The page and the component do not have any difference except for the route generation process. The seperation in this phase isn't justified, but in future realeases this could be helpfull to manage pages and components in seperate ways.
type PageWrapper struct {
	compWrapper *component_generator.ComponentWrapper
	execContext *generator_models.ExecContenxt[FlagsOptions]
}

func NewPageWrapper(
	compWrapper *component_generator.ComponentWrapper,
	execContext *generator_models.ExecContenxt[FlagsOptions],
) *PageWrapper {
	return &PageWrapper{
		compWrapper: compWrapper,
		execContext: execContext,
	}
}

func (w *PageWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	return w.compWrapper.GetDefiners()
}

func (h *PageWrapper) GetEditors() []generator_interfaces.Editor {
	editors := []generator_interfaces.Editor{
		NewAppEditor(h.execContext),
		NewRouteEditor(h.execContext),
	}
	return editors
}
