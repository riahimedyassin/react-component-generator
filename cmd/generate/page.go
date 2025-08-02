package generate

import (
	"os"

	"github.com/riahimedyassin/react-component-generator/config"
	"github.com/riahimedyassin/react-component-generator/config/enums"
	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/lib/files"
	"github.com/riahimedyassin/react-component-generator/pkg/generator"
	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	generator_models "github.com/riahimedyassin/react-component-generator/pkg/generator/models"
	page_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/page"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	route              string // will be used to append the route to the ReactRouter main
	classPageComponent bool
	generatePageCMD    = &cobra.Command{
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
			name = lib.Capitalize(name)
			flags, err := component_generator.NewFlagOptions(cmd)
			if err != nil {
				return err
			}
			pageFlags, err := page_generator.NewFlagsOptions(flags)
			if err != nil {
				return err
			}
			compExecContext := generator_models.NewExecContenxt(name, path, &config.GlobalConfig, *flags)
			pageExecContext := generator_models.NewExecContenxt(name, path, &config.GlobalConfig, *pageFlags)

			fs := files.NewFileSystem()
			compWrapper := component_generator.NewComponentWrapper(compExecContext, fs)
			pageWrapper := page_generator.NewPageWrapper(compWrapper, pageExecContext, fs)
			return generator.NewGenerator().Process(pageWrapper)
		},
	}
)

func setupPageFlagConfig(config *config.Config) {
	generatePageCMD.Flags().StringVarP(&route, "route", "r", "", "Define the page route. This flag is required.")
	generatePageCMD.Flags().BoolVarP(&classPageComponent, "class", "c", config.Component.Type == enums.CLASS, "class component")
	generatePageCMD.MarkFlagRequired("route")
	viper.BindPFlag("route", generatePageCMD.Flags().Lookup("route"))
}

func init() {
	setupPageFlagConfig(&config.GlobalConfig)
}
