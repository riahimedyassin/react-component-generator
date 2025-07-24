package generate

import (
	"fmt"
	"os"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags values.
var (
	compType    string
	withStyling bool
	withTest    bool
)

var (
	generateCompCMD = &cobra.Command{
		Use:   "component",
		Short: "generate a react component",
		Long:  "Generate a react component. For more infos, type in rg g c --help",
		Aliases: []string{
			"c",
			"comp",
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
	}
)

func setupFlags() {
	generateCompCMD.Flags().StringVarP(&compType, "comp", "c", "", "Component type (class or functionnal)")
	generateCompCMD.Flags().BoolVarP(&withStyling, "style", "s", false, "Add styling file to the component")
	generateCompCMD.Flags().BoolVarP(&withTest, "test", "t", false, "Add test file to your component")
	viper.BindPFlag("component.type", generateCompCMD.Flags().Lookup("comp"))
	viper.BindPFlag("component.withTest", generateCompCMD.Flags().Lookup("test"))
	viper.BindPFlag("component.withStyling", generateCompCMD.Flags().Lookup("style"))
}

func init() {
	setupFlags()
	generateCMD.AddCommand(generateCompCMD)
}
