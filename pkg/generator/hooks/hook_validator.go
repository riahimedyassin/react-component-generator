package hooks_generator

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/riahimedyassin/react-component-generator/constants"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

type hookValidator struct {
	execContext     *generator_models.ExecContenxt[FlagsOptions]
	sharedValidator *generator_shared.SharedValidator
}

func newHookValidator(
	execContext *generator_models.ExecContenxt[FlagsOptions],
	sharedValidator *generator_shared.SharedValidator,
) *hookValidator {
	return &hookValidator{
		execContext:     execContext,
		sharedValidator: sharedValidator,
	}
}

func (v *hookValidator) Validate() []error {
	errors := []error{}

	name := v.execContext.Filename

	if err := v.sharedValidator.IsValidateLength(name, 1, constants.FILE_NAME_MAX_LENGTH); err != nil {
		errors = append(errors, err)
	}

	if err := v.sharedValidator.IsAlphaNumeric(name); err != nil {
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
	name := strings.TrimPrefix(v.execContext.Filename, "use")
	return v.sharedValidator.IsValidReactFileFormat(name)
}

func (v *hookValidator) validateReservedNames() error {
	name := strings.ToLower(v.execContext.Filename)

	for _, reserved := range constants.RESERVED_HOOKS_NAMES {
		if name == reserved {
			return fmt.Errorf("hook name '%s' conflicts with a built-in React hook", v.execContext.Filename)
		}
	}

	return nil
}
