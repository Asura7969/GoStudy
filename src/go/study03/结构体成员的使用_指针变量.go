package main

import "fmt"

type Student4 struct {
	id int
	name string
	sex byte // 字符类型
	age int
	address string
}

func main() {
	// 指针有合法指向后才操作成员
	// 先定义一个普通结构体变量
	var s Student4
	// 定义一个指针变量，保存s的地址
	var p1 *Student4
	p1 = &s

	// 通过指针操作成员，p1.id 和 (*p1).id 完全等价,只能使用.运算符
	p1.id = 18
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.address = "sh"
	p1.age = 100
	fmt.Println(p1)

	//通过 new 申请一个结构体
	p2 := new(Student4)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.address = "sh"
	p2.age = 100
	fmt.Println(p2)
}
