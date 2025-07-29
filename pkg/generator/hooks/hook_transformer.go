package hooks_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
)

type hookTransformer struct {
	name, path string
	flags      *FlagsOptions
	config     *config.Config
}

func NewHookTransformer(name, path string,
	flags *FlagsOptions,
	config *config.Config) *hookTransformer {
	return &hookTransformer{
		name:   name,
		path:   path,
		flags:  flags,
		config: config,
	}
}

func (h *hookTransformer) GetTransformed() (*TransformedValues, error) {
	return &TransformedValues{
		HookName: h.transformName(),
	}, nil
}

func (h *hookTransformer) transformName() string {
	res, _ := strings.CutPrefix(h.name, "use")
	res = strings.ToUpper(string(res[0])) + res[1:]
	res = strings.Trim(res, " ")
	res = "use" + res
	return res
}
