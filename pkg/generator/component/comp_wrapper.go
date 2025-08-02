package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

type ComponentWrapper struct {
	execContent *generator_models.ExecContenxt[FlagsOptions]
	fs          *files.FileSystem
}

func NewComponentWrapper(
	execContent *generator_models.ExecContenxt[FlagsOptions],
	fs *files.FileSystem,

) *ComponentWrapper {
	return &ComponentWrapper{
		execContent: execContent,
		fs:          fs,
	}
}

func (c *ComponentWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	definers := []generator_interfaces.FileSpecDefiner{
		newComponentFileSpecGenerator(c.execContent, c.fs),
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

func (h *ComponentWrapper) GetValidators() []generator_interfaces.Validator {
	sharedValidator := generator_shared.NewSharedValidator()
	return []generator_interfaces.Validator{
		newComponentValidator(h.execContent, sharedValidator),
	}
}

func (h *ComponentWrapper) GetCleanUpResult(files []generator_interfaces.GenFiles) error {
	return newComponentCleanUp(h.fs).CleanUp(files)
}
