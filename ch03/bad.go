package ch03

import (
	"fmt"
	"strings"
	"time"
)

func Bad() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	go send(msg)

	for {
		select {
		case m := <-msg:
			fmt.Println(strings.ToUpper(m))
		case <-until:
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan string) {
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}
