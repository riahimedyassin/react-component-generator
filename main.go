package main

import (
	"log"

	"github.com/riahimedyassin/react-component-generator/cmd/root"
)

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
