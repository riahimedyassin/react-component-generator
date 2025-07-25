package generator

import (
	"fmt"
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/config/tokens"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

type ComponentGenerator struct {
	config *config.Config
}

func NewComponentGenerator(config *config.Config) *ComponentGenerator {
	return &ComponentGenerator{
		config: config,
	}
}

// todo :  Check if the name is a path and dive/create the target directory.
func (g *ComponentGenerator) Generate(cwd string, name string) error {
	return g.generateComponent(cwd, name)
}

func (g *ComponentGenerator) generateComponent(cwd, name string) error {
	template, err := g.getTemplate(g.config.Component.Type)
	if err != nil {
		return err
	}
	extraImports, err := g.generateExtraImport(cwd, name)
	if err != nil {
		return err
	}
	populatedTemplate := g.populateTemplate(template, name, extraImports)
	return files.WriteFile(cwd, lib.Capitalize(name), string(g.config.Core.Template), populatedTemplate)
}

func (g *ComponentGenerator) getTemplate(compType enums.ComponentType) (string, error) {
	content, err := files.ReadFile(config.COMP_PATH[compType])
	if err != nil {
		fmt.Print("template not found")
	}
	return string(content), err
}

func (g *ComponentGenerator) populateTemplate(temp string, componentName string, extraImports string) string {
	res := strings.Replace(temp, string(tokens.COMP_NAME), lib.Capitalize(componentName), 1)
	res = strings.Replace(res, string(tokens.EXTRA_IMPORT), extraImports, 1)
	return res
}

func (g *ComponentGenerator) generateExtraImport(cwd, name string) (extraimports string, err error) {
	extraimports = ""
	if g.config.Component.WithStyling {
		extraimports += g.getStylingImport(name)
		if err := g.generateStylingFile(cwd, name); err != nil {
			return "", err
		}
	}
	if g.config.Component.WithTest {
		extraimports += g.getTestingImport(name)
		if err := g.generateTestingFile(cwd, name); err != nil {
			return "", err
		}
	}
	return extraimports, nil
}

func (g *ComponentGenerator) getStylingImport(name string) string {
	return lib.Capitalize(name) + ".css"
}
func (g *ComponentGenerator) getTestingImport(name string) string {
	return lib.Capitalize(name) + ".spec.ts"
}
func (g *ComponentGenerator) generateStylingFile(cwd, name string) error {
	// TODO : dynamic styling extension and support.
	return files.WriteFile(cwd, lib.Capitalize(name), "css", "")
}
func (g *ComponentGenerator) generateTestingFile(cwd, name string) error {
	return files.WriteFile(cwd, lib.Capitalize(name), "spec.ts", "")
}
