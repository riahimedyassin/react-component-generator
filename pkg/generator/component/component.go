package component_generator

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

// TODO : Implement a clean up method if a file creation failed and another succeded
// TODO : Make the generator more generic so that it could handle different types. Maybe a global generator struct will make your code far more DRY.
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	chanErr := make(chan error, 2)
	go func() {
		chanErr <- g.generateComponent()
	}()
	go func() {
		chanErr <- g.generateStyling()
	}()

	var firstErr error
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-chanErr:
			if err != nil && firstErr == nil {
				firstErr = err
			}
		}
	}
	return firstErr
}

func (g *ComponentGenerator) generateComponent() error {
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
	parsedTemplate = strings.Replace(parsedTemplate, string(EXTRA_IMPORT), g.getExtraImport(), 1)
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

func (g *ComponentGenerator) generateStyling() error {
	switch g.config.Project.Styling {
	case enums.NONE, enums.TAILWIND:
		return nil
	}
	if g.config.Component.WithStyling {
		extension := g.getStyleExtension()
		if err := files.WriteFile(g.path, g.name, extension, ""); err != nil {
			return err
		}
	}
	return nil
}

func (g *ComponentGenerator) getStyleExtension() string {
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

func (g *ComponentGenerator) getExtraImport() string {
	res := ""
	if g.config.Component.WithStyling {
		extension := g.getStyleExtension()
		res += "import \"./" + g.name + "." + extension + "\""
	}
	return res
}
