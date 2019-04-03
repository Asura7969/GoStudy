package main

import "fmt"

func main() {

	// iota 常量自动生成器,每个一行,自动累加1
	// iota 给常量复制使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a =%d, b = %d, c =%d\n",a,b,c)

	// iota遇到const , 重置为0
	const d  = iota
	fmt.Printf("d = %d\n",d)

	// 可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d,b1 = %d,c1 = %d\n",a1, b1, c1)

	// 如果是同一行,值都一样
	const (
		u = iota
		m1,m2,m3 = iota,iota,iota
		n = iota
	)
	fmt.Println("=================")
	fmt.Printf("u = %d\n m1 = %d,m2 = %d,m3 = %d\n n = %d\n",u,m1,m2,m3, n)


}
