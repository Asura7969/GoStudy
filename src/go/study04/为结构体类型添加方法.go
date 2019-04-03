package main

import "fmt"

type Person100 struct {
	name string
	sex byte
	id int
}

// 带有接受者的函数叫方法
func (tmp *Person100) PrintlnInfo()  {
	fmt.Println("tmp = ",tmp)
}

// 通过一个函数给成员变量赋值
func (p *Person100)SetInfo(n string,s byte,a int) {
	p.name = n
	p.id = a
	p.sex = s
}

type pointer *int
// pointer是接受者类型，它本身不能是指针类型
//func (tmp pointer)test()  {
//
//}

func main() {
	p := Person100{name:"mike",sex:'b',id:0}
	p.PrintlnInfo()

	// 定义一个结构体变量
	var p2 Person100
	(&p2).SetInfo("yoyo",'b',10)
	p2.PrintlnInfo()


}
