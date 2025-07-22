package config

// Replacable tokens in templates.
type Tokens string

const (
	EXTRA_IMPORT      Tokens = "{{EXTRA_IMPORT}}"
	EXTRA_DEFINITIONS Tokens = "{{EXTRA_DEFINITIONS}}" // Predef values or interfaces
	COMP_NAME         Tokens = "{{COMP_NAME}}"
	DEFAULTED         Tokens = "{{DEFAULTED}}" // Default export.
)
