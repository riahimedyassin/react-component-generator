package config

type ConfigTokens string

const (
	PROJECT_ROUTING_ENTRY_FILE = "project.routing.entryFile"
	PROJECT_BASE               = "project.base"
	PROJECT_STRUCTURE          = "project.structure"
	STYLING                    = "component.styling"
	TYPE                       = "component.type"
	DEFAULT_EXPORT             = "component.defaultExport"
	TEST                       = "component.test"
	WITH_STYLING               = "component.withStyling"
	WITH_TEST                  = "component.withTest"
	WITH_HOOK_TEST             = "hook.withTest"
)
