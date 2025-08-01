package page_generator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/riahimedyassin/react-component-generator/constants"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

type pageValidator struct {
	execContext     *generator_models.ExecContenxt[FlagsOptions]
	sharedValidator *generator_shared.SharedValidator
}

func newPageValidator(
	execContext *generator_models.ExecContenxt[FlagsOptions],
	sharedValidator *generator_shared.SharedValidator,
) *pageValidator {
	return &pageValidator{
		execContext:     execContext,
		sharedValidator: sharedValidator,
	}
}

func (v *pageValidator) Validate() []error {
	errors := []error{}

	if err := v.sharedValidator.IsValidateLength(v.execContext.Filename, 1, constants.FILE_NAME_MAX_LENGTH); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateName(); err != nil {
		errors = append(errors, err)
	}

	if err := v.sharedValidator.IsValidReactFileFormat(v.execContext.Filename); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateRoute(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (v *pageValidator) validateName() error {
	name := v.execContext.Filename

	if err := v.sharedValidator.IsAlphaNumeric(name); err != nil {
		return err
	}

	return v.sharedValidator.IsPascalCase(name)
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
