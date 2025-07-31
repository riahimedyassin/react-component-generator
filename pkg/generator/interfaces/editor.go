package generator_interfaces

import generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"

// an editor should implement this interface.
type Editor interface {
	GetEditFileSpec() (*generator_models.EditFileSpec, error)
}
