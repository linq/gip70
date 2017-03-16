package ch02

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

func Ic() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
