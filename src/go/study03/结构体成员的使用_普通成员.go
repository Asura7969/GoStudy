package main

import "fmt"

type Student3 struct {
	id int
	name string
	sex byte // 字符类型
	age int
	address string
}
func main() {
	// 定义一个结构体普通变量
	var s Student3
	s.id = 1
	s.name = "a"
	s.sex = 'w'
	s.age = 10
	s.address = "sh"
	fmt.Println(s)
}
