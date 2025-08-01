package hooks_generator

import (
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookValidator struct {
	execContext *generator_models.ExecContenxt[FlagsOptions]
}

func newHookValidator(
	execContext *generator_models.ExecContenxt[FlagsOptions],
) *hookValidator {
	return &hookValidator{
		execContext: execContext,
	}
}

func (v *hookValidator) Validate() error {

	return nil
}
