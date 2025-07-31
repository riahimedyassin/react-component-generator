package config

import "github.com/riahimedyassin/react-component-generator/lib"

type TemplatePaths string

var (
	CLASS_COMPONENT_TEMPLATE_PATH = lib.GetTemplatePath("comp.class.tpl")
	FUNC_COMPONENT_TEMPLATE_PATH  = lib.GetTemplatePath("comp.func.tpl")
	HOOK_TEMPLATE_PATH            = lib.GetTemplatePath("hook.tpl")
	HOOK_TEST_TEMPLATE_PATH       = lib.GetTemplatePath("hook.test.tpl")
	APP_TEMPLATE_PATH             = lib.GetTemplatePath("app.tpl")
	PROJECT_CONFIG_PATH           = ".rcg/rg.config.json"
	ROUTE_CONFIG_PATH             = ".rcg/rg.route.json"
)
