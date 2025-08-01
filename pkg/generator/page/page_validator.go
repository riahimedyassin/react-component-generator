package page_generator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type pageValidator struct {
	execContext   *generator_models.ExecContenxt[FlagsOptions]
	maxNameLength int
}

func newPageValidator(
	execContext *generator_models.ExecContenxt[FlagsOptions],
) *pageValidator {
	return &pageValidator{
		execContext:   execContext,
		maxNameLength: 80,
	}
}

func (v *pageValidator) Validate() []error {
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

	if err := v.validateRoute(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (v *pageValidator) validateLength() error {
	length := len(v.execContext.Filename)
	if length > v.maxNameLength {
		return fmt.Errorf("page name too long: %d characters, enter a shorter page name", length)
	}
	if length == 0 {
		return fmt.Errorf("page name cannot be empty")
	}
	return nil
}

func (v *pageValidator) validateName() error {
	name := v.execContext.Filename

	// Check if name contains only valid characters (alphanumeric, underscore)
	validNameRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	if !validNameRegex.MatchString(name) {
		return fmt.Errorf("page name '%s' contains invalid characters. Use only letters, numbers, and underscores, starting with a letter", name)
	}

	// Check if first character is uppercase (PascalCase convention)
	if !unicode.IsUpper(rune(name[0])) {
		return fmt.Errorf("page name '%s' should start with an uppercase letter (PascalCase convention)", name)
	}

	return nil
}

func (v *pageValidator) validateFormat() error {
	name := v.execContext.Filename

	// Check for consecutive underscores
	if strings.Contains(name, "__") {
		return fmt.Errorf("page name '%s' should not contain consecutive underscores", name)
	}

	// Check if name ends with underscore
	if strings.HasSuffix(name, "_") {
		return fmt.Errorf("page name '%s' should not end with underscore", name)
	}

	// Check for numbers immediately after the first character
	if len(name) > 1 && unicode.IsDigit(rune(name[1])) {
		return fmt.Errorf("page name '%s' should not have numbers immediately after the first letter", name)
	}

	// Check for all uppercase names (should be PascalCase, not SCREAMING_SNAKE_CASE)
	if strings.ToUpper(name) == name && len(name) > 1 {
		return fmt.Errorf("page name '%s' should use PascalCase, not all uppercase", name)
	}

	// Check for camelCase instead of PascalCase
	if len(name) > 0 && unicode.IsLower(rune(name[0])) {
		return fmt.Errorf("page name '%s' should use PascalCase (start with uppercase letter)", name)
	}

	return nil
}

func (v *pageValidator) validateRoute() error {
	route := v.execContext.Flags.Route

	if route == "" {
		return fmt.Errorf("route cannot be empty")
	}

	// Route should start with /
	if !strings.HasPrefix(route, "/") {
		return fmt.Errorf("route '%s' should start with '/'", route)
	}

	// Route should not end with / unless it's the root route
	if route != "/" && strings.HasSuffix(route, "/") {
		return fmt.Errorf("route '%s' should not end with '/' unless it's the root route", route)
	}

	// Check for valid route characters
	validRouteRegex := regexp.MustCompile(`^/[a-zA-Z0-9/_:-]*$`)
	if !validRouteRegex.MatchString(route) {
		return fmt.Errorf("route '%s' contains invalid characters. Use only letters, numbers, hyphens, underscores, colons, and forward slashes", route)
	}

	// Check for consecutive slashes
	if strings.Contains(route, "//") {
		return fmt.Errorf("route '%s' should not contain consecutive slashes", route)
	}

	// Check route depth (too many nested levels might be problematic)
	routeParts := strings.Split(strings.Trim(route, "/"), "/")
	if len(routeParts) > 5 && route != "/" {
		return fmt.Errorf("route '%s' is too deeply nested (%d levels). Consider simplifying the route structure", route, len(routeParts))
	}

	return nil
}
