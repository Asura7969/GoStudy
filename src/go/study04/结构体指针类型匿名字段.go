package main

import "fmt"

type Person03 struct{
	name string
	sex byte
	age int
}

type Student03 struct{
	*Person03  // 指针类型
	id int
	address string
	name string
}

func main() {
	s1 := Student03{Person03:&Person03{"mike",'m',18},id:666,address:"bj",name:"kk"}
	fmt.Println(s1.Person03)
	fmt.Println(s1.name,s1.sex,s1.age,s1.id,s1.address)

	// 先定义变量
	var s2 Student03
	s2.Person03 = new(Person03) // 分配空间
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 52
	s2.id = 1
	s2.address = "cz"
	fmt.Println(s2.name,s2.sex,s2.age,s2.id,s2.address)
}

