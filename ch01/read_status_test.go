package ch01

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestStatus(t *testing.T) {
	conn, _ := net.Dial("tcp", "example.com:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
