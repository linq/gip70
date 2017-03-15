package ch02

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var ops struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func GfParse() {
	flags.Parse(&ops)
	if ops.Spanish == true {
		fmt.Printf("Hola %s!\n", ops.Name)
	} else {
		fmt.Printf("Hollo %s!\n", ops.Name)
	}
}
