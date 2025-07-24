package generate

import (
	"github.com/riahimedyassin/react-component-generator/cmd/root"
	"github.com/spf13/cobra"
)

var (
	generateCMD = &cobra.Command{
		Use:   "generate",
		Short: "generate everything related to react.",
		Long:  "The generate command is used to generate your components, pages, contexts, hooks etc.\nThe generate command should be followed by one of the upcoming commands:\n* p : Generate page. Type in rg g --help for more.",
		Aliases: []string{
			"g",
			"gen",
		},
	}
)

func init() {
	root.RootCmd.AddCommand(generateCMD)
}
