package main

import "fmt"

func main() {
	flag := true

	fmt.Printf("flag = %t \n", flag)

	// bool类型不能转化为int
	//fmt.Printf("flag = %d \n", int(flag))

	var char byte
	char = 'a'
	var t int
	t = int(char)
	fmt.Printf("t = ",t)

}
