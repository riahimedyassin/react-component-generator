package page_generator

import (
	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
)

// Responsible for editing app tsx or the routing file.
type PageEditor struct {
	name, path string
	config     *config.Config
	flags      any
}

func NewPageEditor() *PageEditor {
	return &PageEditor{}
}

func (e *PageEditor) GetEditSpecs() (*generator.EditFileSpec, error) {
	return nil, nil
}
