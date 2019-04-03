package main

import "fmt"

type Humanger interface {
	// 方法，只有声明，没有实现，由别的类型（自定义类型实现）
	sayHi()
}

type S1 struct {
	name string
}

func (tmp *S1)sayHi() {
	fmt.Printf("s1[%s]\n",tmp.name)
}

type S2 struct {
	name string
}

func (tmp S2)sayHi() {
	fmt.Printf("s2[%s]\n",tmp.name)
}

func WhoSayHi(i Humanger) {
	i.sayHi()
}
// 多态的表现
func main() {
	s1 := S1{"S1"}
	s2 := S2{"S2"}
	// 调用同一函数,不同表现,多态,多种形态
	WhoSayHi(&s1)
	WhoSayHi(s2)
}

// 接口的定义和实现
func main0() {
	// 定义接口类型变量
	var i Humanger

	// 只是实现了此接口,那么这个类型的变量（接受者类型）就可以给i赋值
	s1 := S1{"S1"}
	i = &s1
	i.sayHi()

	s2 := S2{"S2"}
	i = s2
	i.sayHi()
}