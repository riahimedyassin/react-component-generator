package page_generator

type RouteElement struct {
	Path       string `json:"path"`
	Component  string `json:"component"`
	ImportPath string `json:"importPath"`
}
