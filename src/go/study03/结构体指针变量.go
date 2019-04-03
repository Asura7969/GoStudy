package main

import "fmt"

type Student2 struct {
	id int
	name string
	sex byte // 字符类型
	age int
	address string
}

func main() {
	// 顺序初始化,每个成员必须初始化
	var p *Student2 = &Student2{1,"lisi",'m',20,"sh"}
	fmt.Println(p)

	// 指定成员初始化,没有指定的成员,自动赋值
	s2 := &Student2{name:"mike",address:"sh"}
	fmt.Printf("type s2 is:%T\n",s2)
	fmt.Println(s2)
}

