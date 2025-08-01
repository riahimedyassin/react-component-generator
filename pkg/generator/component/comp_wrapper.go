package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type ComponentWrapper struct {
	execContent *generator_models.ExecContenxt[FlagsOptions]
}

func NewComponentWrapper(
	execContent *generator_models.ExecContenxt[FlagsOptions],
) *ComponentWrapper {
	return &ComponentWrapper{
		execContent: execContent,
	}
}

func (c *ComponentWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	definers := []generator_interfaces.FileSpecDefiner{
		newComponentFileSpecGenerator(c.execContent, files.NewFileSystem()),
	}
	if c.execContent.Config.Component.WithStyling {
		definers = append(definers, newStyleGenerator(c.execContent))
	}
	if c.execContent.Config.Component.WithTest {
		definers = append(definers, newTestGenerator(c.execContent))
	}
	return definers
}

func (h *ComponentWrapper) GetEditors() []generator_interfaces.Editor {
	return nil
}
