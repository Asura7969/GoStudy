package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte // 字符类型
	age int
	address string
}

func main() {
	// 顺序初始化,每个成员必须初始化
	var s1 Student = Student{1,"lisi",'m',20,"sh"}
	fmt.Println(s1)

	// 指定成员初始化,没有指定的成员,自动赋值
	s2 := Student{name:"mike",address:"sh"}
	fmt.Println(s2)
}
