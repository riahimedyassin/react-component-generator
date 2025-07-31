package page_generator

import (
	"fmt"
	"strings"

	component_generator "github.com/riahimedyassin/react-component-generator/pkg/generator/component"
	"github.com/spf13/viper"
)

type FlagsOptions struct {
	ClassComponent bool
	Route          string
}

func NewFlagsOptions(compFlags *component_generator.FlagsOptions) (*FlagsOptions, error) {
	route := viper.GetString("route")
	if strings.Trim(route, " ") == "" {
		return nil, fmt.Errorf("provide a route")
	}
	return &FlagsOptions{
		ClassComponent: compFlags.ClassComponent,
		Route:          route,
	}, nil
}
