package ch01

import (
	"fmt"
	"testing"
)

func Names2() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func TestName2(t *testing.T) {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
