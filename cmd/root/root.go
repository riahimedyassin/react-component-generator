package root

import (
	"fmt"

	"github.com/riahimedyassin/react-component-generator/cmd/generate"
	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command{
		Use:   "rg",
		Short: "React Generator - V0.0.1",
		Long:  "React Generator is a pacakge used to manage your react project seamlisly with ease",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Hello world")
		},
		Aliases: []string{
			"reactgenerator",
		},
	}
)

func init() {
	RootCmd.AddCommand(generate.GenerateCMD)

}

func Execute() error {
	return RootCmd.Execute()
}
