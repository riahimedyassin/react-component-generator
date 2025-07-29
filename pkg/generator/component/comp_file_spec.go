package component_generator

import (
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

type componentFileSpecGenerator struct {
	name, path string
	config     *config.Config
	flags      *FlagsOptions
	fs         *files.FileSystem
}

func newComponentFileSpecGenerator(name, path string, config *config.Config, options *FlagsOptions, fs *files.FileSystem) *componentFileSpecGenerator {
	return &componentFileSpecGenerator{
		name:   name,
		config: config,
		flags:  options,
		path:   path,
		fs:     fs,
	}
}

func (c *componentFileSpecGenerator) GetFileSpec() (*generator.FileSpec, error) {
	template, err := c.getTemplate()
	if err != nil {
		return nil, err
	}
	content, err := c.generateContent(template)
	if err != nil {
		return nil, err
	}
	return generator.NewFileSpec(c.name, c.path, c.getExtension(), content), nil
}

func (c *componentFileSpecGenerator) generateContent(template string) (string, error) {
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

func (g *componentFileSpecGenerator) getExtraDefinitions() string {
	res := ""
	if g.config.Project.Base == enums.TYPESCRIPT {
		res += "interface " + g.name + "Props {\n}\n"
	}
	return res
}

func (g *componentFileSpecGenerator) getExtension() string {
	if g.config.Project.Base == enums.JAVASCRIPT {
		return "jsx"
	}
	return "tsx"
}

func (g *componentFileSpecGenerator) getTemplate() (string, error) {
	var path string
	if g.flags.ClassComponent || g.config.Component.Type == enums.CLASS {
		path = config.CLASS_COMPONENT_TEMPLATE_PATH
	} else {
		path = config.FUNC_COMPONENT_TEMPLATE_PATH
	}
	content, err := g.fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (g *componentFileSpecGenerator) getStyleExtension() string {
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

func (g *componentFileSpecGenerator) getExtraImport() string {
	res := ""
	if g.config.Component.WithStyling {
		extension := g.getStyleExtension()
		res += "import \"./" + g.name + "." + extension + "\""
	}
	return res
}
