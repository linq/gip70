package ch03

import (
	"fmt"
	"time"
)

func CloseSend() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send2(ch)

	for {
		select {
		case st := <-ch:
			fmt.Println("Got message.", st)
		case <-timeout:
			fmt.Println("time out", timeout)
			return
		default:
			fmt.Println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send2(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	fmt.Println("Sent and closed")
}
