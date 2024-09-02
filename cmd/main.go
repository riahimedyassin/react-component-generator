package main

import (
	"fmt"

	"github.com/riahimedyassin/react-component-generator/lib"
	"github.com/riahimedyassin/react-component-generator/types"
)

func main() {
	fmt.Println("Hello world")
	content, _ := lib.ReadFile("./", "config.json")
	var config types.Config
	lib.ParseJson(content, &config)
}
