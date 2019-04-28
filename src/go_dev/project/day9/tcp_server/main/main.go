package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("start server")
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("listener failed, err:", err)
		return
	}

	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}

		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	for true {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}

		data := string(buf[0:n])
		for _, v := range data {
			fmt.Printf("%d", v)
		}

		fmt.Println("read:", strings.TrimSpace(string(buf)))
	}
}
