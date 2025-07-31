package hooks_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookFileSpecGenerator struct {
	name    string
	path    string
	config  *config.Config
	options *FlagsOptions
	fs      *files.FileSystem
}

func newHookFileSpecGenerator(
	name string,
	path string,
	config *config.Config,
	options *FlagsOptions,
	fs *files.FileSystem,
) *hookFileSpecGenerator {
	return &hookFileSpecGenerator{
		name:    name,
		path:    path,
		config:  config,
		options: options,
		fs:      fs,
	}
}

func (g *hookFileSpecGenerator) GetFileSpec() (*generator_models.FileSpec, error) {
	template, err := g.getTemplate()
	if err != nil {
		return nil, err
	}
	content, err := g.generateContent(template)
	if err != nil {
		return nil, err
	}
	fileSpec := &generator_models.FileSpec{
		Name:      g.name,
		Content:   content,
		Path:      g.path,
		Extension: g.getExtension(),
	}
	return fileSpec, nil
}

func (g *hookFileSpecGenerator) generateContent(template string) (string, error) {
	return strings.Replace(template, string(HOOK_NAME), g.name, 1), nil
}

func (g *hookFileSpecGenerator) getTemplate() (string, error) {
	template, err := g.fs.ReadFile(config.HOOK_TEMPLATE_PATH)
	if err != nil {
		return "", err
	}
	return string(template), nil
}

func (g *hookFileSpecGenerator) getExtension() string {
	if g.config.Project.Base == enums.JAVASCRIPT {
		return "js"
	}
	return "ts"
}

// todo : implement
/*
refineHookName takes a given name and refactor it if needed based on those constraints :
- Should start with use (auth -> useAuth)
- First letter after use if upper case (useauth -> useAuth)
- should be a valid name ()
*/
// func refineHookName(name string) {

// }
