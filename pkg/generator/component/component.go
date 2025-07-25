package component_generator

import (
	"path/filepath"
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

type ComponentGenerator struct {
	path    string
	name    string
	config  *config.Config
	options *FlagsOptions
}

// Full file path will include the file name and the relative directories something like /features/user/Hello
func NewComponentGenerator(fullFilePath string, config *config.Config, options *FlagsOptions) *ComponentGenerator {
	name := filepath.Base(fullFilePath)
	path := filepath.Dir(fullFilePath)
	return &ComponentGenerator{
		config:  config,
		options: options,
		name:    lib.Capitalize(name),
		path:    path,
	}
}

// todo :  Check if the name is a path and dive/create the target directory.
func (g *ComponentGenerator) Generate() error {
	template, err := g.getTemplate()
	if err != nil {
		return err
	}
	parsedTemplate := g.parseTemplate(template)
	if err := files.WriteFile(g.path, g.name, g.getExtension(), parsedTemplate); err != nil {
		return err
	}
	return nil
}

func (g *ComponentGenerator) parseTemplate(template string) string {
	var parsedTemplate string
	defaultedExport := "" // default export a component or nomral export.
	params := ""
	if g.config.Component.DefaultExport {
		defaultedExport = "default"
	}
	if g.config.Project.Base == enums.TYPESCRIPT {
		params = "{params} : " + g.name + "Props"
	}
	parsedTemplate = strings.Replace(template, string(EXTRA_DEFINITIONS), g.getExtraDefinitions(), 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(EXTRA_IMPORT), "", 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(DEFAULTED), defaultedExport, 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(COMP_NAME), g.name, 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(EXTRA_PARAMS), params, 1)
	return parsedTemplate
}

func (g *ComponentGenerator) getExtension() string {
	if g.config.Project.Base == enums.JAVASCRIPT {
		return "jsx"
	}
	return "tsx"
}

func (g *ComponentGenerator) getExtraDefinitions() string {
	res := ""
	if g.config.Project.Base == enums.TYPESCRIPT {
		res += "interface " + g.name + "Props {\n}\n"
	}
	return res
}

func (g *ComponentGenerator) getTemplate() (string, error) {
	var path string
	if g.options.ClassComponent || g.config.Component.Type == enums.CLASS {
		path = config.CLASS_COMPONENT_TEMPLATE_PATH
	} else {
		path = config.FUNC_COMPONENT_TEMPLATE_PATH
	}
	content, err := files.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// func (g *ComponentGenerator) generateStyling() error {
// 	extensions := map[enums.Styling]string{
// 		enums.SCSS: "scss",
// 		enums.CSS:  "css",
// 	}
// 	switch g.config.Project.Styling {
// 	case enums.NONE, enums.TAILWIND:
// 		return nil
// 	}
// 	if g.config.Component.WithStyling {
// 		// validation will be handeled at an upper level, not to worry about here.
// 		// ! validation will happen at an upper level in an upcoming validation package.
// 		extension := extensions[g.config.Project.Styling]
// 	}
// 	return nil
// }
