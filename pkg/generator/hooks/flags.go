package hooks_generator

import "github.com/spf13/cobra"

type FlagsOptions struct {
	WithTests bool
}

func NewFlagOptions(cmd *cobra.Command) (*FlagsOptions, error) {
	withTests, err := cmd.Flags().GetBool("t")
	if err != nil {
		return nil, err
	}
	return &FlagsOptions{
		WithTests: withTests,
	}, nil
}
