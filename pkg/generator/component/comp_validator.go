package component_generator

import (
	"github.com/riahimedyassin/react-component-generator/constants"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

type componentValidator struct {
	execContext     *generator_models.ExecContenxt[FlagsOptions]
	sharedValidator *generator_shared.SharedValidator
}

func newComponentValidator(execContext *generator_models.ExecContenxt[FlagsOptions], sharedValidator *generator_shared.SharedValidator) *componentValidator {
	return &componentValidator{
		execContext:     execContext,
		sharedValidator: sharedValidator,
	}
}

func (v *componentValidator) Validate() []error {
	errors := []error{}

	name := v.execContext.Filename

	if err := v.sharedValidator.IsValidateLength(name, 1, constants.FILE_NAME_MAX_LENGTH); err != nil {
		errors = append(errors, err)
	}

	if err := v.validateName(); err != nil {
		errors = append(errors, err)
	}

	if err := v.sharedValidator.IsValidReactFileFormat(name); err != nil {
		errors = append(errors, err)
	}

	return errors
}

func (v *componentValidator) validateName() error {
	name := v.execContext.Filename

	if err := v.sharedValidator.IsAlphaNumeric(name); err != nil {
		return err
	}
	return v.sharedValidator.IsPascalCase(name)
}
