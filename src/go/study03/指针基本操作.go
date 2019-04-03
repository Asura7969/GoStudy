package main

import "fmt"

func main() {
	a := 10

	fmt.Printf("a = %d\n",a)
	fmt.Printf("a = %v\n",&a)

	// 保存某个变量的地址，需要指针类型  *int 保存int的地址， **int 保存 *int 地址
	// 什么是定义? 定义只是特殊的声明
	// 定义一个变量,类型为 *int
	var p *int // 指针变量指向谁,就把谁的地址赋值给指针变量
	p = &a
	fmt.Printf("p = %v,&a = %v\n",p,&a)

	*p = 666 // *p 操作的不是p的内存，是p所指向的内存(就是a)
	fmt.Printf("*p = %v,a = %v\n",*p,a)
	fmt.Printf("p = %d\n",p)
}
