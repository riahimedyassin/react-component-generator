package hooks_generator

import (
	"errors"

	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type hookCleanUp struct {
	fs *files.FileSystem
}

func newHookCleanUp(fs *files.FileSystem) *hookCleanUp {
	return &hookCleanUp{
		fs: fs,
	}
}

func (c *hookCleanUp) CleanUp(files []generator_interfaces.GenFiles) error {
	errs := []error{}
	for _, file := range files {
		if fileEdit, ok := file.(*generator_models.FileSpec); ok {
			if err := c.cleanUpFileSpec(*fileEdit); err != nil {
				errs = append(errs, err)
			}
		}
		if editFile, ok := file.(*generator_models.EditFileSpec); ok {
			if err := c.cleanUpEditFileSpec(*editFile); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.Join(errs...)
}

func (c *hookCleanUp) cleanUpFileSpec(file generator_models.FileSpec) error {
	return c.fs.DeleteFile(file.Path, file.Name, file.Extension)
}

func (c *hookCleanUp) cleanUpEditFileSpec(file generator_models.EditFileSpec) error {
	return c.fs.WriteFile(file.Path, file.Name, file.Extension, file.CurrentContent)
}
