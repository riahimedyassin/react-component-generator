package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags values.
var (
	class       bool
	withStyling bool
	withTest    bool
)

var (
	generateCompCMD = &cobra.Command{
		Use:   "c  [component name]",
		Short: "generate a react component",
		Long:  "Generate a react component. For more infos, type in rg g c --help",
		Aliases: []string{
			"component",
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
			name, path := lib.Capitalize(filepath.Base(args[0])), filepath.Dir(args[0])
			flagsOptions, err := component_generator.NewFlagOptions(cmd)
			if err != nil {
				return err
			}
			compFileSpec, err := component_generator.
				NewComponentFileSpecGenerator(name,
					path,
					&config.GlobalConfig,
					flagsOptions,
					component_generator.NewStyleGenerator(name, path, &config.GlobalConfig, flagsOptions),
				).
				GetFileSpec()
			if err != nil {
				return err
			}
			return generator.NewGenerator().Generate(*compFileSpec)
			// compGen := component_generator.NewComponentGenerator(args[0], &config.GlobalConfig, flagsOptions)
			// return compGen.Generate()
		},
	}
)

func setupFlags() {
	generateCompCMD.Flags().BoolVarP(&class, "class", "c", false, "Class component")
	generateCompCMD.Flags().BoolVarP(&withStyling, "style", "s", false, "Add styling file to the component")
	generateCompCMD.Flags().BoolVarP(&withTest, "test", "t", false, "Add test file to your component")
	// VIPER BINDING
	viper.BindPFlag(config.TEST, generateCompCMD.Flags().Lookup("test"))
	viper.BindPFlag(config.WITH_STYLING, generateCompCMD.Flags().Lookup("style"))
}

func init() {
	setupFlags()
}
