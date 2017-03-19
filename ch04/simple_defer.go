package ch04

import (
	"fmt"
)

func SimpleDefer() {
	defer goodbye()

	fmt.Println("Hello world.")
}

func goodbye() {
	fmt.Println("Goodbye")
}
