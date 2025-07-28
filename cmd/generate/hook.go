package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	hooks_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/hooks"
	"github.com/spf13/cobra"
)

var generateHookCMD = &cobra.Command{
	Use:   "h [name]",
	Short: "generata a react hook",
	Aliases: []string{
		"hook",
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you passed in %d argument, enter only one valid name", len(args))
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		return config.Load(cwd)
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		name, path := filepath.Base(args[0]), filepath.Dir(args[0])
		flags, err := hooks_generator.NewFlagOptions(cmd)
		if err != nil {
			return err
		}
		hookWrapper := hooks_generator.NewHookWrapper(name, path, &config.GlobalConfig, flags)
		return generator.NewGenerator().Generate(hookWrapper)
	},
}
