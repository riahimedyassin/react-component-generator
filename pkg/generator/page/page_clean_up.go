package page_generator

import (
	"errors"

	"github.com/riahimedyassin/react-component-generator/lib/files"
	generator_interfaces "github.com/riahimedyassin/react-component-generator/pkg/generator/interfaces"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
)

type pageCleanUp struct {
	fs *files.FileSystem
}

func newPageCleanUp(fs *files.FileSystem) *pageCleanUp {
	return &pageCleanUp{
		fs: fs,
	}
}

func (c *pageCleanUp) CleanUp(files []generator_interfaces.GenFiles) error {
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

func (c *pageCleanUp) cleanUpFileSpec(file generator_models.FileSpec) error {
	return c.fs.DeleteFile(file.Path, file.Name, file.Extension)
}

func (c *pageCleanUp) cleanUpEditFileSpec(file generator_models.EditFileSpec) error {
	return c.fs.WriteFile(file.Path, file.Name, file.Extension, file.CurrentContent)
}
