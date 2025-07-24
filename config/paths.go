package config

import "github.com/riahimedyassin/react-component-generator/config/enums"

type TemplatePaths string

const (
	CLASS_COMPONENT = "./templates/comp.class.tmp"
	FUNC_COMPONENT  = "./templates/comp.func.tmp"
)

var COMP_PATH = map[enums.ComponentType]string{
	enums.CLASS:       CLASS_COMPONENT,
	enums.FUNCTIONNAL: FUNC_COMPONENT,
}
