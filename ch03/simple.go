package ch03

import (
	"fmt"
	"runtime"
)

func Simple() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again.")
	runtime.Gosched()
}
