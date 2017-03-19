package ch03

import (
	"time"
)

func TSleep() {
	time.Sleep(5 * time.Second)

	sleep := time.After(5 * time.Second)
	<-sleep
}
