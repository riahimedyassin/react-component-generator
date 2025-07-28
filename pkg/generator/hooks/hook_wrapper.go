package hooks_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type HookWrapper struct {
	name, path string
	config     *config.Config
	flags      *FlagsOptions
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
	}
}

func (h *HookWrapper) GetDefiners() []generator.FileSpecDefiner {
	definers := []generator.FileSpecDefiner{
		newHookFileSpecGenerator(h.name, h.path, h.config, h.flags),
	}
	return definers
}
