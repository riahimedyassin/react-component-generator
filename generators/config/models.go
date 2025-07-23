package config_generator

type GenerateConfigFlags struct {
	Typescript bool
	Linting    bool
	Prettier   bool
}

func GenerateConfigFlagsFromString(execFlags string) *GenerateConfigFlags {
	config := &GenerateConfigFlags{}

	return config
}
