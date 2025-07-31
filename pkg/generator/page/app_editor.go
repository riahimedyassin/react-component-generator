package page_generator

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

// Responsible for editing app tsx or the routing file.
type AppEditor struct {
	pageName, pagePath string
	appName            string // The routing tsx file, by default app.tsx
	config             *config.Config
	flags              *FlagsOptions
	fs                 *files.FileSystem
}

func NewAppEditor(pageName, pagePath string,
	config *config.Config,
	flags *FlagsOptions) *AppEditor {
	return &AppEditor{
		appName:  config.Project.Routing.EntryFile,
		pageName: pageName,
		pagePath: pagePath,
		config:   config,
		flags:    flags,
		fs:       files.NewFileSystem(),
	}
}

func (e *AppEditor) GetEditFileSpec() (*generator_models.EditFileSpec, error) {
	routes, err := e.getRouteConfig()
	if err != nil {
		return nil, err
	}
	content, err := e.getAppContent(routes)
	if err != nil {
		return nil, err
	}
	snapshot, err := e.getAppSnapshot()
	if err != nil {
		return nil, err
	}
	return &generator_models.EditFileSpec{
		Name:           e.appName,
		Path:           "./",
		NewContent:     content,
		Extension:      fmt.Sprintf("%sx", e.config.GetProjectFileExtension()),
		CurrentContent: snapshot,
	}, nil

}

// getRouteConfig retrieves the routing file related to the CLI.
// the file is a JSON file.
func (e *AppEditor) getRouteConfig() ([]RouteElement, error) {
	content, err := e.fs.ReadFile(config.ROUTE_CONFIG_PATH)
	if err != nil {
		return nil, err
	}
	var routeGroup []RouteElement
	if err := json.Unmarshal([]byte(content), &routeGroup); err != nil {
		return nil, err
	}
	// Appending the current route being generated.
	routeGroup = append(routeGroup, e.getRouteElement())
	return routeGroup, err
}

// the end route file generation content after compiling
func (e *AppEditor) getAppContent(routes []RouteElement) (string, error) {
	template, err := e.getAppTemplate()
	if err != nil {
		return "", err
	}
	res := strings.Replace(template, ROUTES, e.getAppRoutes(routes), 1)
	res = strings.Replace(res, ROUTES, e.getAppImports(routes), 1)
	return res, nil
}

func (e *AppEditor) getAppTemplate() (string, error) {
	content, err := e.fs.ReadFile(config.APP_TEMPLATE_PATH)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (e *AppEditor) getAppImports(routes []RouteElement) string {
	res := ""
	for _, route := range routes {
		res += "import " + route.ImportPath + "\n"
	}
	return res
}

func (e *AppEditor) getAppRoutes(routes []RouteElement) string {
	res := ""
	for _, route := range routes {
		res += "<Route path=\" " + route.Path + "\" element={<" + route.Component + "/>} />"
	}
	return res
}

func (e *AppEditor) getAppSnapshot() (string, error) {
	content, err := e.fs.ReadFile(e.appName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (e *AppEditor) getRouteElement() RouteElement {
	importPath := filepath.Join(e.pagePath, e.pageName)
	return RouteElement{
		Path:       e.flags.Route,
		Component:  e.pageName,
		ImportPath: importPath,
	}
}
