package component_generator

import "github.com/spf13/cobra"

// The flag options override some default config values.
// In this case, if a class component flag is determined, it will override the default config component type.
type FlagsOptions struct {
	ClassComponent bool
}

func NewFlagOptions(cmd *cobra.Command) (*FlagsOptions, error) {
	classFlag, err := cmd.Flags().GetBool("class")
	if err != nil {
		return nil, err
	}
	return &FlagsOptions{
		ClassComponent: classFlag,
	}, nil
}
