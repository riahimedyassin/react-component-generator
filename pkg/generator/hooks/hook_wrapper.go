package hooks_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
)

type HookWrapper struct {
	name, path string
	config     *config.Config
	flags      *FlagsOptions
	fs         *files.FileSystem
}

func NewHookWrapper(
	name, path string,
	config *config.Config,
	flags *FlagsOptions,
) *HookWrapper {
	return &HookWrapper{
		name:   name,
		path:   path,
		config: config,
		flags:  flags,
		fs:     files.NewFileSystem(),
	}
}

func (h *HookWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	fs := files.NewFileSystem()
	definers := []generator_interfaces.FileSpecDefiner{
		newHookFileSpecGenerator(h.name, h.path, h.config, h.flags, fs),
	}
	if h.flags.WithTests {
		definers = append(definers, newHookTestGenerator(h.name, h.path, h.config, h.flags, fs))
	}
	return definers
}

func (h *HookWrapper) GetEditors() []generator_interfaces.Editor {
	return nil
}
