package generator

import (
	"context"
	"sync"

	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
)

// Responsible for file generation
type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(filesSpecs ...FileSpec) error {
	doneFiles := lib.NewSafeSlice[FileSpec]() // Tracking the generated files to safely cleanup in case of an error.
	errChan := make(chan error, len(filesSpecs))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	for _, f := range filesSpecs {
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
		}(f)
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
