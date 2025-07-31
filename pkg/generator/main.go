package generator

import (
	"context"
	"errors"
	"sync"
	"time"

	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

// Responsible for file generation
type Generator struct {
	fs *files.FileSystem
}

func NewGenerator() *Generator {
	return &Generator{
		fs: files.NewFileSystem(),
	}
}

// Entry point.
func (g *Generator) Process(wrapper generator_interfaces.Wrapper) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	errChan := make(chan error, 10) // Buffered channel
	var pwg sync.WaitGroup

	pwg.Add(2)
	go func() {
		defer pwg.Done()
		if err := g.edit(ctx, errChan, wrapper); err != nil {
			errChan <- err
			cancel()
		}
	}()

	go func() {
		defer pwg.Done()
		if err := g.generate(ctx, errChan, wrapper); err != nil {
			errChan <- err
			cancel()
		}
	}()

	pwg.Wait()
	cancel()
	close(errChan) // Close channel after all goroutines complete

	// Collect all errors
	if err := <-errChan; err != nil {
		return err
	}
	return nil
}
func (g *Generator) edit(pctx context.Context, errChan chan error, wrapper generator_interfaces.Wrapper) error {
	editors := wrapper.GetEditors()
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	for _, editor := range editors {
		editFilesSpecs, err := editor.GetEditFileSpec()
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(e *generator_models.EditFileSpec) {
			select {
			case <-ctx.Done():
				return
			default:
				defer wg.Done()
				if err := g.fs.WriteFile(e.Path, e.Name, e.Extension, e.NewContent); err != nil {
					errChan <- err
					cancel()
					return
				}
			}

		}(editFilesSpecs)
	}
	wg.Wait()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (g *Generator) generate(pctx context.Context, errChan chan error, wrapper generator_interfaces.Wrapper) error {
	fileSpecDefs := wrapper.GetDefiners()
	doneFiles := lib.NewSafeSlice[generator_models.FileSpec]() // Tracking the generated files to safely cleanup in case of an error.
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	var wg sync.WaitGroup
	for _, f := range fileSpecDefs {
		// todo : move in withing a goroutine
		fileSpec, err := f.GetFileSpec()
		if err != nil {
			if errors.Is(err, &rg_errors.IgnoreDefinerError{}) {
				continue
			}
			// call for clean up func here
			return err
		}
		wg.Add(1)
		go func(fileSpec generator_models.FileSpec) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				if err := g.fs.WriteFile(fileSpec.Path, fileSpec.Name, fileSpec.Extension, fileSpec.Content); err != nil {
					errChan <- err
					cancel()
					return
				}
				doneFiles.Append(fileSpec)
			}
		}(*fileSpec)
	}
	wg.Wait()
	select {
	case err := <-errChan:
		g.cleanUp(doneFiles.GetAll())
		return err
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

// TODO : Implement
// in case of an error the function will rollback the created files to keep a clean project structure
func (g *Generator) cleanUp(files []generator_models.FileSpec) {

}
