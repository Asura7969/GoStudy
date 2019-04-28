package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, e := net.Dial("tcp", "0.0.0.0:9999")
	if e != nil {
		fmt.Println("Error dialing", e.Error())
		return
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for true {
		input, _ := inputReader.ReadString('\n')
		trim := strings.Trim(input, "\r\n")
		if trim == "q" {
			return
		}
		_, e := conn.Write([]byte(trim))
		if e != nil {
			return
		}
	}
}
