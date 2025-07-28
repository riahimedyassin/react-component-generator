package generator

import (
	"context"
	"errors"
	"sync"

	rg_errors "github.com/riahimedyassin/react-component-generator/errors"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

// Responsible for file generation
type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(filesSpecs ...FileSpecDefiner) error {
	doneFiles := lib.NewSafeSlice[FileSpec]() // Tracking the generated files to safely cleanup in case of an error.
	errChan := make(chan error, len(filesSpecs))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	for _, f := range filesSpecs {
		fileSpec, err := f.GetFileSpec()
		if err != nil {
			if errors.Is(err, &rg_errors.IgnoreDefinerError{}) {
				continue
			}
			// call for clean up func here
			return err
		}
		wg.Add(1)
		go func(fileSpec FileSpec) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				if err := files.WriteFile(fileSpec.Path, fileSpec.Name, fileSpec.Extension, fileSpec.Content); err != nil {
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
	default:
		return nil
	}
}

// TODO : Implement
// in case of an error the function will rollback the created files to keep a clean project structure
func (g *Generator) cleanUp(files []FileSpec) {

}
