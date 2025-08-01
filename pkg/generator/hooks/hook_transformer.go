package hooks_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookTransformer struct {
	name, path string
	flags      *FlagsOptions
	config     *config.Config
}

func NewHookTransformer(execContext *generator_models.ExecContenxt[FlagsOptions]) *hookTransformer {
	return &hookTransformer{
		name:   execContext.Filename,
		path:   execContext.Filepath,
		flags:  &execContext.Flags,
		config: execContext.Config,
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
