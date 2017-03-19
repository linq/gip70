package ch03

import (
	"fmt"
	"os"
	"sync"
)

func WaitGz() {
	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}
