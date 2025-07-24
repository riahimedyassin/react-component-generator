package generate

import "github.com/spf13/cobra"

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

func init() {
	generatePageCMD.Flags().StringVarP(&route, "route", "r", "", "Define the page route. This flag is required.")
	generatePageCMD.MarkFlagRequired("route")
	generateCMD.AddCommand(generateCMD)
}
