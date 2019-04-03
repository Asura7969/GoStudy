package main

import "fmt"

type Psson struct {
	name string
	age int
}

func (p Psson)SetInfoValue() {
	fmt.Printf("SetInfoValue:%p,%v\n",&p,p)
}

func (p *Psson)SetInfoPointer() {
	fmt.Printf("SetInfoPointer:%p,%v\n",p,p)
}

// 方法值
func main02() {
	p := Psson{"mike",1}
	fmt.Printf("main:%p,%v\n",&p,p)

	// 传统方式调用
	p.SetInfoPointer()

	// 保存方式入口地址
	pFunc := p.SetInfoPointer // 这个就是方法值,调用函数时,无需再传递接受者,隐藏了接受者
	pFunc()

	vFunc := p.SetInfoValue
	vFunc()
}
// 方法表达式
func main() {
	p := Psson{"mike",1}
	fmt.Printf("main:%p,%v\n",&p,p)

	// f := p.SetInfoValue() // 隐藏了接受者
	// 方法表达式
	f:=(*Psson).SetInfoPointer
	f(&p) // 显示把接受者传递过去 ===> p.SetInfoPointer()

	f2 := (Psson).SetInfoValue
	f2(p) // 显示把接受者传递过去 ===> p.SetInfoValue()
}