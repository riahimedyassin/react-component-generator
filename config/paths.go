package config

import (
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib"
)

type TemplatePaths string

var (
	CLASS_COMPONENT = lib.GetTemplatePath("comp.class.tmp")
	FUNC_COMPONENT  = lib.GetTemplatePath("comp.func.tmp")
)

var COMP_PATH = map[enums.ComponentType]string{
	enums.CLASS:      CLASS_COMPONENT,
	enums.FUNCTIONAL: FUNC_COMPONENT,
}
