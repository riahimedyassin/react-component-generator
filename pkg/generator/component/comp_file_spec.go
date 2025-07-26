package component_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type ComponentFileSpecGenerator struct {
	name           string
	path           string
	config         *config.Config
	options        *FlagsOptions
	styleGenerator *StyleGenerator
}

func NewComponentFileSpecGenerator(name, path string, config *config.Config, options *FlagsOptions, styleGenerator *StyleGenerator) *ComponentFileSpecGenerator {
	return &ComponentFileSpecGenerator{
		name:           name,
		config:         config,
		options:        options,
		path:           path,
		styleGenerator: styleGenerator,
	}
}

func (c *ComponentFileSpecGenerator) GetFileSpec() (*generator.FileSpec, error) {
	content, err := c.generateContent()
	if err != nil {
		return nil, err
	}
	if err := c.styleGenerator.Generate(); err != nil {
		return nil, err
	}
	return generator.NewFileSpec(c.name, c.path, c.getExtension(), content), nil
}

func (c *ComponentFileSpecGenerator) generateContent() (string, error) {
	template, err := c.getTemplate()
	if err != nil {
		return "", err
	}
	var parsedTemplate string
	defaultedExport := "" // default export a component or nomral export.
	params := ""
	if c.config.Component.DefaultExport {
		defaultedExport = "default"
	}
	if c.config.Project.Base == enums.TYPESCRIPT {
		params = "{params} : " + c.name + "Props"
	}
	parsedTemplate = strings.Replace(template, string(EXTRA_DEFINITIONS), c.getExtraDefinitions(), 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(EXTRA_IMPORT), c.getExtraImport(), 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(DEFAULTED), defaultedExport, 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(COMP_NAME), c.name, 1)
	parsedTemplate = strings.Replace(parsedTemplate, string(EXTRA_PARAMS), params, 1)
	return parsedTemplate, nil
}

func (g *ComponentFileSpecGenerator) getExtraDefinitions() string {
	res := ""
	if g.config.Project.Base == enums.TYPESCRIPT {
		res += "interface " + g.name + "Props {\n}\n"
	}
	return res
}

func (g *ComponentFileSpecGenerator) getExtension() string {
	if g.config.Project.Base == enums.JAVASCRIPT {
		return "jsx"
	}
	return "tsx"
}

func (g *ComponentFileSpecGenerator) getTemplate() (string, error) {
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

func (g *ComponentFileSpecGenerator) getStyleExtension() string {
	extensions := map[enums.Styling]string{
		enums.SCSS: "scss",
		enums.CSS:  "css",
	}
	extension, ok := extensions[g.config.Project.Styling]
	if !ok {
		extension = "css" // default style
	}
	return extension
}

func (g *ComponentFileSpecGenerator) getExtraImport() string {
	res := ""
	if g.config.Component.WithStyling {
		extension := g.getStyleExtension()
		res += "import \"./" + g.name + "." + extension + "\""
	}
	return res
}
