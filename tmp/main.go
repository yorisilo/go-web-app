package main

import (
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:1111")
	defer listener.Close()

	conn, err := listener.Accept()
	defer conn.Close()

	buf := make([]byte, 4*1024)

	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
	}

}
