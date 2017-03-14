package ch01

import (
	"fmt"
	"net/http"
	"testing"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func TestHello(t *testing.T) {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}
