package hooks_generator

import (
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

type HookWrapper struct {
	execContext *generator_models.ExecContenxt[FlagsOptions]
	fs          *files.FileSystem
}

func NewHookWrapper(
	execContext *generator_models.ExecContenxt[FlagsOptions],
) *HookWrapper {
	return &HookWrapper{
		execContext: execContext,
		fs:          files.NewFileSystem(),
	}
}

func (h *HookWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	fs := files.NewFileSystem()
	definers := []generator_interfaces.FileSpecDefiner{
		newHookFileSpecGenerator(h.execContext, fs),
	}
	if h.execContext.Flags.WithTests {
		definers = append(definers, newHookTestGenerator(h.execContext, fs))
	}
	return definers
}

func (h *HookWrapper) GetEditors() []generator_interfaces.Editor {
	return nil
}

func (h *HookWrapper) GetValidators() []generator_interfaces.Validator {
	sharedValidator := generator_shared.NewSharedValidator()
	return []generator_interfaces.Validator{
		newHookValidator(h.execContext, sharedValidator),
	}
}
