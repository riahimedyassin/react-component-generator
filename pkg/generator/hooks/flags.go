package hooks_generator

import "github.com/spf13/cobra"

type FlagsOptions struct {
}

func NewFlagOptions(cmd *cobra.Command) (*FlagsOptions, error) {
	return &FlagsOptions{}, nil
}
