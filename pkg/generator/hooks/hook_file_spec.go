package hooks_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type HookFileSpecGenerator struct {
	name     string
	path     string
	config   *config.Config
	options  *FlagsOptions
	testFile *HookTestGenerator
}

func NewHookFileSpecGenerator(
	name string,
	path string,
	config *config.Config,
	options *FlagsOptions,
	testFile *HookTestGenerator,
) *HookFileSpecGenerator {
	return &HookFileSpecGenerator{
		name:     name,
		path:     path,
		config:   config,
		options:  options,
		testFile: testFile,
	}
}

func (g *HookFileSpecGenerator) GetFileSpec() (*generator.FileSpec, error) {
	content, err := g.generateContent()
	if err != nil {
		return nil, err
	}
	if err := g.testFile.Generate(); err != nil {
		return nil, err
	}
	fileSpec := &generator.FileSpec{
		Name:      g.name,
		Content:   content,
		Path:      g.path,
		Extension: g.getExtension(),
	}
	return fileSpec, nil
}

func (g *HookFileSpecGenerator) generateContent() (string, error) {
	res := ""
	template, err := g.getTemplate()
	if err != nil {
		return "", err
	}
	res = strings.Replace(template, string(HOOK_NAME), g.name, 1)
	return res, nil
}

func (g *HookFileSpecGenerator) getTemplate() (string, error) {
	template, err := files.ReadFile(config.HOOK_TEMPLATE_PATH)
	if err != nil {
		return "", err
	}
	return string(template), nil
}

func (g *HookFileSpecGenerator) getExtension() string {
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
func refineHookName(name string) {

}
