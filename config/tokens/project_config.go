package tokens

type TProjectLintingTokens string

const (
	ESLINT   TProjectLintingTokens = "{{ESLINT}}"
	PRETTIER TProjectLintingTokens = "{{PRETTIER}}"
)

type TProjectCoreTokens string

const (
	STRUCTURE TProjectCoreTokens = "{{STRUCTURE}}"
	TEMPLATE  TProjectCoreTokens = "{{TEMPLATE}}"
)

type TProjectStateTokens string

const (
	STATE TProjectStateTokens = "{{STATE}}" // type of the state (redux, zustand etc)
)

type TProjectComponentTokens string

const ()

const (
	STYLING     TProjectComponentTokens = "{{STYLING}}"     // type of the style
	WITHSTYLING TProjectComponentTokens = "{{WITHSTYLING}}" // include styles?
	WITHTEST    TProjectComponentTokens = "{{WITHTEST}}"
	TYPE        TProjectComponentTokens = "{{TYPE}}" // class of func component.
)
