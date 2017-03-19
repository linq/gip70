package ch03

import (
	"fmt"
	"strings"
	"time"
)

func Good() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	go send3(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(strings.ToUpper(m))
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send3(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
