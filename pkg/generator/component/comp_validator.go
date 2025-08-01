package component_generator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type componentValidator struct {
	execContext   *generator_models.ExecContenxt[FlagsOptions]
	maxNameLength int
}

func newComponentValidator(execContext *generator_models.ExecContenxt[FlagsOptions]) *componentValidator {
	return &componentValidator{
		execContext:   execContext,
		maxNameLength: 80,
	}
}

func (v *componentValidator) Validate() []error {
	errors := []error{}

	if err := v.validateLength(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateName(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateFormat(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (v *componentValidator) validateLength() error {
	length := len(v.execContext.Filename)
	if length > v.maxNameLength {
		return fmt.Errorf("component name too long: %d characters, enter a shorter component name", length)
	}
	if length == 0 {
		return fmt.Errorf("component name cannot be empty")
	}
	return nil
}

func (v *componentValidator) validateName() error {
	name := v.execContext.Filename

	// Check if name contains only valid characters (alphanumeric, underscore)
	validNameRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	if !validNameRegex.MatchString(name) {
		return fmt.Errorf("component name '%s' contains invalid characters. Use only letters, numbers, and underscores, starting with a letter", name)
	}

	// Check if first character is uppercase (PascalCase convention)
	if !unicode.IsUpper(rune(name[0])) {
		return fmt.Errorf("component name '%s' should start with an uppercase letter (PascalCase convention)", name)
	}

	return nil
}

func (v *componentValidator) validateFormat() error {
	name := v.execContext.Filename

	// Check for consecutive underscores
	if strings.Contains(name, "__") {
		return fmt.Errorf("component name '%s' should not contain consecutive underscores", name)
	}

	// Check if name ends with underscore
	if strings.HasSuffix(name, "_") {
		return fmt.Errorf("component name '%s' should not end with underscore", name)
	}

	// Check for numbers immediately after the first character
	if len(name) > 1 && unicode.IsDigit(rune(name[1])) {
		return fmt.Errorf("component name '%s' should not have numbers immediately after the first letter", name)
	}

	// Check for all uppercase names (should be PascalCase, not SCREAMING_SNAKE_CASE)
	if strings.ToUpper(name) == name && len(name) > 1 {
		return fmt.Errorf("component name '%s' should use PascalCase, not all uppercase", name)
	}

	// Check for camelCase instead of PascalCase
	if len(name) > 0 && unicode.IsLower(rune(name[0])) {
		return fmt.Errorf("component name '%s' should use PascalCase (start with uppercase letter)", name)
	}

	return nil
}
