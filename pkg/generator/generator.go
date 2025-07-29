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
	fs *files.FileSystem
}

func NewGenerator() *Generator {
	return &Generator{
		fs: files.NewFileSystem(),
	}
}

func (g *Generator) Generate(wrapper Wrapper) error {
	fileSpecDefs := wrapper.GetDefiners()
	doneFiles := lib.NewSafeSlice[FileSpec]() // Tracking the generated files to safely cleanup in case of an error.
	errChan := make(chan error, len(fileSpecDefs))
	ctx, cancel := context.WithCancel(context.Background())
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
		go func(fileSpec FileSpec) {
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
	default:
		return nil
	}
}

// TODO : Implement
// in case of an error the function will rollback the created files to keep a clean project structure
func (g *Generator) cleanUp(files []FileSpec) {

}
