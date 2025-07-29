package generate

import (
	"os"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	page_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/page"
	"github.com/spf13/cobra"
)

var (
	route           string // will be used to append the route to the ReactRouter main
	generatePageCMD = &cobra.Command{
		Use:   "p",
		Short: "generate a page",
		Long:  "Generate a react page. The page will be appended to the react router tree. Therefor a -r flag is required to generate corretly the page.",
		Aliases: []string{
			"page",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			return config.Load(cwd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			name, path := lib.GetNameAndPath(args[0])
			flags, err := component_generator.NewFlagOptions(cmd)
			if err != nil {
				return err
			}
			compWrapper := component_generator.NewComponentWrapper(name, path, &config.GlobalConfig, flags)
			pageWrapper := page_generator.NewPageWrapper(compWrapper)
			return generator.NewGenerator().Generate(pageWrapper)
		},
	}
)

func init() {
	generatePageCMD.Flags().StringVarP(&route, "route", "r", "", "Define the page route. This flag is required.")
	generatePageCMD.MarkFlagRequired("route")
}
