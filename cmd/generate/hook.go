package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	hooks_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/hooks"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
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
		execContent := generator_models.NewExecContenxt(name, path, &config.GlobalConfig, *flags)

		// Transformation phase
		transformer := hooks_generator.NewHookTransformer(execContent)
		transformedValues, err := transformer.GetTransformed()
		if err != nil {
			return err
		}
		execContent.Filename = transformedValues.HookName
		hookWrapper := hooks_generator.NewHookWrapper(execContent)
		return generator.NewGenerator().Process(hookWrapper)
	},
}

func init() {
	generateHookCMD.Flags().BoolP("test", "t", config.GlobalConfig.Hooks.WithTest, "include tests")
}
