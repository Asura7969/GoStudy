package main

import "fmt"

// 函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int{
	var x int

	return func() int {
		x ++
		return x * x
	}
}
// 普通函数
func test001() int {
	var x int
	x ++
	return x * x
}

func main() {

	fmt.Println(test001())
	fmt.Println(test001())
	fmt.Println(test001())
	fmt.Println(test001())
	fmt.Println(test001())

	fmt.Println("调用闭包函数=============")
	// 返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	// 它不关心这些捕获的变量和常量是否已经超出了作用域
	// 所以只有闭包还在使用它，这些变量就还会存在
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}


