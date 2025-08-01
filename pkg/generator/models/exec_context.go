package generator_models

import "github.com/riahimedyassin/react-component-generator/config"

type ExecContenxt[T any] struct {
	Filename, Filepath string
	Config             *config.Config
	Flags              T
}

func NewExecContenxt[T any](
	Filename, Filepath string,
	Config *config.Config,
	Flags T,
) *ExecContenxt[T] {
	return &ExecContenxt[T]{
		Filename: Filename,
		Filepath: Filepath,
		Config:   Config,
		Flags:    Flags,
	}
}
