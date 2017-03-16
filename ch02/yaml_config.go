package ch02

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func Yc() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
