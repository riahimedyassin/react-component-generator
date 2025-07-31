package page_generator

import (
	"encoding/json"
	"path/filepath"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

// Responsible for editing the route config file
type RouteEditor struct {
	name, path string
	config     *config.Config
	flags      *FlagsOptions
	fs         *files.FileSystem
}

func NewRouteEditor(
	name, path string,
	config *config.Config,
	flags *FlagsOptions,
) *RouteEditor {
	return &RouteEditor{
		name:   name,
		path:   path,
		config: config,
		flags:  flags,
	}
}

func (e *RouteEditor) GetEditFileSpec() (*generator_models.EditFileSpec, error) {
	routes, err := e.getRouteConfig()
	if err != nil {
		return nil, err
	}
	content, err := e.getRouteContent(routes)
	if err != nil {
		return nil, err
	}
	snapshot, err := e.getRouteSnaphot()
	if err != nil {
		return nil, err
	}
	return generator_models.NewEditFileSpec("rg.route", ".rcg", "json", content, snapshot), nil
}

func (e *RouteEditor) getRouteConfig() ([]RouteElement, error) {
	content, err := e.fs.ReadFile(config.ROUTE_CONFIG_PATH)
	if err != nil {
		return nil, err
	}
	var routeGroup []RouteElement
	if err := json.Unmarshal([]byte(content), &routeGroup); err != nil {
		return nil, err
	}
	return routeGroup, err
}

// Writeable end content after editing.
func (e *RouteEditor) getRouteContent(routes []RouteElement) (string, error) {
	routes = append(routes, e.getRouteElement())
	marshalledRoutes, err := json.Marshal(routes)
	if err != nil {
		return "", err
	}
	return string(marshalledRoutes), nil
}

func (e *RouteEditor) getRouteElement() RouteElement {
	importPath := filepath.Join(e.path, e.name)
	return RouteElement{
		Path:       e.flags.Route,
		Component:  e.name,
		ImportPath: importPath,
	}
}

func (e *RouteEditor) getRouteSnaphot() (string, error) {
	content, err := e.fs.ReadFile(config.ROUTE_CONFIG_PATH)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
