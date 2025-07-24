package tokens

type CoreConfigTokens = string

const (
	STRUCTURE CoreConfigTokens = "core.structure"
	TEMPLATE  CoreConfigTokens = "core.template"
	STYLING   CoreConfigTokens = "core.styling"
)

type ComponentConfigTokens = string

const (
	WITH_STYLING ComponentConfigTokens = "component.withStyling"
	WITH_TEST    ComponentConfigTokens = "component.withTest"
	TYPE         ComponentConfigTokens = "component.type"
)

type StateConfigTokens = string

const (
	STATE_TYPE StateConfigTokens = "state.type"
)

type LintingConfigTokens = string

const (
	ESLINT   LintingConfigTokens = "linting.eslint"
	PRETTIER LintingConfigTokens = "linting.prettier"
)
