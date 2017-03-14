package ch01

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp, _ := http.Get("http://example.com/")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
