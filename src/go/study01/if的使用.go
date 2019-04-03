package main

import "fmt"

func main() {
	s := "aaa"

	if s == "www" {
		fmt.Println(true)
	}
	
	// if 支持一个初始化语句
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a != 10")
	}
}
