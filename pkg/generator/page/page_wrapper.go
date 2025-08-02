package page_generator

import (
	"sync"

	"github.com/riahimedyassin/react-component-generator/lib/files"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	generator_shared "github.com/riahimedyassin/react-component-generator/pkg/generator/shared"
)

// The page and the component do not have any difference except for the route generation process. The seperation in this phase isn't justified, but in future realeases this could be helpfull to manage pages and components in seperate ways.
type PageWrapper struct {
	compWrapper *component_generator.ComponentWrapper
	execContext *generator_models.ExecContenxt[FlagsOptions]
	fs          *files.FileSystem
}

func NewPageWrapper(
	compWrapper *component_generator.ComponentWrapper,
	execContext *generator_models.ExecContenxt[FlagsOptions],
	fs *files.FileSystem,
) *PageWrapper {
	return &PageWrapper{
		compWrapper: compWrapper,
		execContext: execContext,
		fs:          fs,
	}
}

func (w *PageWrapper) GetDefiners() []generator_interfaces.FileSpecDefiner {
	return w.compWrapper.GetDefiners()
}

func (h *PageWrapper) GetEditors() []generator_interfaces.Editor {
	editors := []generator_interfaces.Editor{
		NewAppEditor(h.execContext),
		NewRouteEditor(h.execContext),
	}
	return editors
}

func (h *PageWrapper) GetValidators() []generator_interfaces.Validator {
	compValidators := h.compWrapper.GetValidators()
	sharedValidator := generator_shared.NewSharedValidator()
	return append(compValidators, newPageValidator(h.execContext, sharedValidator))
}

func (h *PageWrapper) GetCleanUpResult(files []generator_interfaces.GenFiles) error {
	errChan := make(chan error, len(files))
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		errChan <- h.compWrapper.GetCleanUpResult(files)
	}()

	go func() {
		defer wg.Done()
		errChan <- newPageCleanUp(h.fs).CleanUp(files)
	}()

	wg.Wait()
	return <-errChan
}
