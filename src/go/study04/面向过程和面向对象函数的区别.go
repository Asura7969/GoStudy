package main

import "fmt"

// 面向过程
func Add01(a, b int) int {
	return a + b
}

// 面向对象，方法：给某个类型绑定一个函数
type testvar int
func (t testvar)Add02(b testvar) testvar {
	return t + b
}
func main() {
	var result int
	result = Add01(1,1)
	fmt.Println(result)

	var a testvar = 2// 先定义一个变量
	r := a.Add02(4)
	fmt.Println(r)
}
