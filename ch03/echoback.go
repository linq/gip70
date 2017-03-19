package ch03

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Atask() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Time out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
