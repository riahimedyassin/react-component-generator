package generator_interfaces

import generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"

type FileSpecDefiner interface {
	GetFileSpec() (*generator_models.FileSpec, error)
}
