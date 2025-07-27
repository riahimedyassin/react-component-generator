package generate

import (
	"os"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/spf13/cobra"
)

var (
	GenerateCMD = &cobra.Command{
		Use:   "g",
		Short: "generate everything related to react.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			return config.Load(cwd)
		},
		Long: "The generate command is used to generate your components, pages, contexts, hooks etc.\nThe generate command should be followed by one of the upcoming commands:\n* p : Generate page. Type in rg g --help for more.",
		Aliases: []string{
			"generate",
		},
	}
)

func init() {
	GenerateCMD.AddCommand(generateCompCMD)
	GenerateCMD.AddCommand(generatePageCMD)
	GenerateCMD.AddCommand(generateHookCMD)
}
