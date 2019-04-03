package main

import "fmt"

type Person02 struct{
	name string
	sex byte
	age int
}

type Student02 struct{
	Person02
	id int
	address string
	name string
}


func main() {
	var s Student02
	s.name = "mike" // 如果在作用域中能找到调用的字段就给该字段赋值, 如果找不到,到父级里面去找
	fmt.Println("s = ",s)

	// 显示调用Person02中的name字段
	s.Person02.name = "mike02"
	fmt.Println("s = ",s)


}
