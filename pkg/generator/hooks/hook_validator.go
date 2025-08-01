package hooks_generator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookValidator struct {
	execContext   *generator_models.ExecContenxt[FlagsOptions]
	maxNameLength int
}

func newHookValidator(
	execContext *generator_models.ExecContenxt[FlagsOptions],
) *hookValidator {
	return &hookValidator{
		execContext:   execContext,
		maxNameLength: 80,
	}
}

func (v *hookValidator) Validate() []error {
	errors := []error{}

	if err := v.validateLength(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateName(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validatePrefix(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateFormat(); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateReservedNames(); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (v *hookValidator) validateLength() error {
	length := len(v.execContext.Filename)
	if length > v.maxNameLength {
		return fmt.Errorf("%d characters, enter a shorter hook name", len(v.execContext.Filename))
	}
	if length == 0 {
		return fmt.Errorf("hook name cannot be empty")
	}
	return nil
}
func (v *hookValidator) validateName() error {
	name := v.execContext.Filename

	// Check if name contains only valid characters (alphanumeric, underscore)
	validNameRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	if !validNameRegex.MatchString(name) {
		return fmt.Errorf("hook name '%s' contains invalid characters. Use only letters, numbers, and underscores, starting with a letter", name)
	}

	return nil
}

func (v *hookValidator) validatePrefix() error {
	name := v.execContext.Filename

	if !strings.HasPrefix(name, "use") {
		return fmt.Errorf("hook name '%s' must start with 'use' prefix", name)
	}

	if name == "use" {
		return fmt.Errorf("hook name cannot be just 'use', it needs a descriptive name after 'use'")
	}

	// Check if the character after "use" is uppercase (following React conventions)
	if len(name) > 3 && !unicode.IsUpper(rune(name[3])) {
		return fmt.Errorf("hook name '%s' should follow camelCase convention (e.g., 'useAuth', 'useState')", name)
	}

	return nil
}

func (v *hookValidator) validateFormat() error {
	name := v.execContext.Filename

	// Check for consecutive underscores
	if strings.Contains(name, "__") {
		return fmt.Errorf("hook name '%s' should not contain consecutive underscores", name)
	}

	// Check if name ends with underscore
	if strings.HasSuffix(name, "_") {
		return fmt.Errorf("hook name '%s' should not end with underscore", name)
	}

	// Check for numbers immediately after "use"
	if len(name) > 3 && unicode.IsDigit(rune(name[3])) {
		return fmt.Errorf("hook name '%s' should not have numbers immediately after 'use'", name)
	}

	return nil
}

func (v *hookValidator) validateReservedNames() error {
	name := strings.ToLower(v.execContext.Filename)

	// List of React built-in hooks and common reserved names
	reservedNames := []string{
		"usestate", "useeffect", "usecontext", "usereducer", "usecallback",
		"usememo", "useref", "useimperativehandle", "uselayouteffect",
		"usedebugvalue", "usedeferredvalue", "useid", "useinsertioneffect",
		"usesyncexternalstore", "usetransition", "useactionstate",
		"useformstatus", "useoptimistic",
	}

	for _, reserved := range reservedNames {
		if name == reserved {
			return fmt.Errorf("hook name '%s' conflicts with a built-in React hook", v.execContext.Filename)
		}
	}

	return nil
}
