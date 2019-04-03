package main

import "fmt"

type Class struct {
	id int
}
type myStr string

type room struct {
	Class // 结构体匿名字段
	int // 非结构体匿名字段
	myStr
}

	func main() {
		s := room{Class:Class{1},int:1,myStr:"a"}
		fmt.Println(s)

}
