package main

import "fmt"

type Person6 struct{
	name string
	id int
	sex byte
}

// 接受者为普通变量,值传递
func (p Person6)setValue(name string, id int, sex byte) {
	p.sex = sex
	p.id = id
	p.name = name
	fmt.Printf("setValue &p = %p\n",&p)
}

// 接受者为指针变量,引用传递
func (p *Person6)setPoint(name string, id int, sex byte) {
	p.sex = sex
	p.id = id
	p.name = name
	fmt.Printf("setPoint &p = %p\n",&p)
}

func main() {
	p := Person6{name:"go",id:11,sex:'b'}
	fmt.Printf("main &p = %p\n",&p)
	p.setValue("java",100,'j')
	//p.setPoint("flink",3,'f')
	fmt.Printf("main &p = %p\n",&p)
	fmt.Println("p = ",p)
}