package main

import "fmt"

type Pers struct {
	name string
	id int
	sex byte
}

func (p Pers)setValue() {
	fmt.Println("setValue")
}

func (p *Pers)setPoint() {
	fmt.Println("setPoint")
}
// 普通变量
func main() {
	p := Pers{"go",1,'v'}
	p.setPoint() // 内部先把p,转化为&p.setPoint()
}


// 指针变量
func main01() {
	// 结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集合
	p := &Pers{"go",1,'v'}
	p.setPoint() // func (p *Pers)setPoint()
	(*p).setPoint() //把(*p)转换成p后再调用，等于上面

	// 内部的转化,先把指针p,转化成*p后再调用
	(*p).setValue()
	p.setValue()

}