package generator_shared

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type SharedValidator struct {
}

func NewSharedValidator() *SharedValidator {
	return &SharedValidator{}
}

func (s *SharedValidator) IsValidateLength(value string, min, max int) error {
	length := len(value)
	if length > max {
		return fmt.Errorf("%s too long: %d characters, enter a shorter value", value, length)
	}
	if length <= min {
		return fmt.Errorf("%s too short: %d characters, enter a longer value", value, length)
	}
	return nil
}

func (s *SharedValidator) IsPascalCase(value string) error {
	if !unicode.IsUpper(rune(value[0])) {
		return fmt.Errorf("'%s' should start with an uppercase letter (PascalCase convention)", value)
	}
	return nil

}

// Check if name contains only valid characters (alphanumeric, underscore)
func (s *SharedValidator) IsAlphaNumeric(value string) error {
	validNameRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	if !validNameRegex.MatchString(value) {
		return fmt.Errorf("page name '%s' contains invalid characters. Use only letters, numbers, and underscores, starting with a letter", value)
	}
	return nil
}

func (s *SharedValidator) IsValidReactFileFormat(value string) error {
	// Check for consecutive underscores
	if strings.Contains(value, "__") {
		return fmt.Errorf("value '%s' should not contain consecutive underscores", value)
	}

	// Check if value ends with underscore
	if strings.HasSuffix(value, "_") {
		return fmt.Errorf("value '%s' should not end with underscore", value)
	}

	// Check for numbers immediately after the first character
	if len(value) > 1 && unicode.IsDigit(rune(value[1])) {
		return fmt.Errorf("value '%s' should not have numbers immediately after the first letter", value)
	}

	// Check for all uppercase values (should be PascalCase, not SCREAMING_SNAKE_CASE)
	// todo : skip this once a transformer is implemented.
	if strings.ToUpper(value) == value && len(value) > 1 {
		return fmt.Errorf("value '%s' should use PascalCase, not all uppercase", value)
	}

	// Check for camelCase instead of PascalCase
	// todo : skip this once a transformer is implemented.
	if len(value) > 0 && unicode.IsLower(rune(value[0])) {
		return fmt.Errorf("value '%s' should use PascalCase (start with uppercase letter)", value)
	}

	return nil
}
