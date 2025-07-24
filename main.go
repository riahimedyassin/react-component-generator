package main

import (
	"fmt"
	"log"
	"os"

	"github.com/riahimedyassin/react-component-generator/cmd"
	"github.com/riahimedyassin/react-component-generator/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Cannot use CLI , %v", err)
		return
	}
	_, err = config.Load(cwd)
	if err != nil {
		log.Fatal(err.Error())
	}
	cmd.Execute()
}
