package ch01

import (
	"fmt"
	"testing"
)

func Names() (string, string) {
	return "Foo", "Bar"
}

func TestReturn(t *testing.T) {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	n3, _ := Names()
	fmt.Println(n3)
}
