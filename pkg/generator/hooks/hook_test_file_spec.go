package hooks_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookTestGenerator struct {
	name   string
	path   string
	config *config.Config
	flags  *FlagsOptions
	fs     *files.FileSystem
}

func newHookTestGenerator(execContext *generator_models.ExecContenxt[FlagsOptions], fs *files.FileSystem) *hookTestGenerator {
	return &hookTestGenerator{
		name:   execContext.Filename,
		path:   execContext.Filepath,
		config: execContext.Config,
		flags:  &execContext.Flags,
		fs:     fs,
	}
}

func (g *hookTestGenerator) GetFileSpec() (*generator_models.FileSpec, error) {
	if g.config.Component.WithTest {
		template, err := g.getTemplate()
		if err != nil {
			return nil, err
		}
		content := g.generateContent(template)
		return generator_models.NewFileSpec(g.name, g.path, g.getExtension(), content), nil
	}
	return nil, rg_errors.NewIgnoreDefinerError()
}

func (g *hookTestGenerator) getExtension() string {
	if g.config.Project.Base == enums.TYPESCRIPT {
		return "spec.ts"
	}
	return "spec.js"
}

func (g *hookTestGenerator) getTemplate() (string, error) {
	content, err := g.fs.ReadFile(config.HOOK_TEST_TEMPLATE_PATH)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (g *hookTestGenerator) generateContent(template string) string {
	return strings.ReplaceAll(template, string(HOOK_NAME), g.name)
}
