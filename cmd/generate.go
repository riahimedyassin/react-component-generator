package cmd

import "github.com/spf13/cobra"

var (
	generateCMD = &cobra.Command{
		Use:   "generate",
		Short: "generate everything related to react.",
		Long:  "The generate command is used to generate your components, pages, contexts, hooks etc.\nThe generate command should be followed by one of the upcoming commands:\n* p : Generate page. Type in g p --help for more.",
	}
)

var (
	route           string // will be used to append the route to the ReactRouter main
	generatePageCMD = &cobra.Command{
		Use:   "page",
		Short: "generate a page",
		Long:  "Generate a react page. The page will be appended to the react router tree. Therefor a -r flag is required to generate corretly the page.",
		Aliases: []string{
			"p",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func setupGeneratePageCMD() {
	generatePageCMD.Flags().StringVarP(&route, "route", "r", "", "Define the page route. This flag is required.")
	generatePageCMD.MarkFlagRequired("route")
}

func init() {
	setupGeneratePageCMD()
	generateCMD.AddCommand(generatePageCMD)
	rootCmd.AddCommand(generateCMD)
}
