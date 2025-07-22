package main

import (
	"fmt"

	"github.com/riahimedyassin/react-component-generator/config"
)

func main() {
	config := config.NewConfig("./")
	if err := config.Load(); err != nil {
		fmt.Printf("Cannot load config , %v", err)
		return
	}
}
