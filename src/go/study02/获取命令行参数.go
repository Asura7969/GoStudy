package main

import (
	"fmt"
	"os"
)

func main() {
	array := os.Args

	l := len(array)

	fmt.Println("参数长度：",l)

	for i, data := range array {
		fmt.Printf("array[%d],%s\n",i,data)
	}
}
