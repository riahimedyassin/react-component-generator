package tokens

// Replacable tokens in templates.
type TComponentTokens string

const (
	EXTRA_IMPORT      TComponentTokens = "{{EXTRA_IMPORT}}"
	EXTRA_DEFINITIONS TComponentTokens = "{{EXTRA_DEFINITIONS}}" // Predef values or interfaces
	COMP_NAME         TComponentTokens = "{{COMP_NAME}}"
	DEFAULTED         TComponentTokens = "{{DEFAULTED}}" // Default export.
)
