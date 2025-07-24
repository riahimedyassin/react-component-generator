package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command{
		Use:   "reactgenerator",
		Short: "React Generator - V0.0.1",
		Long:  "React Generator is a pacakge used to manage your react project seamlisly with ease",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Hello world")
		},
		Aliases: []string{
			"rg",
		},
	}
)

func init() {

}

func Execute() {
	RootCmd.Execute()
}
