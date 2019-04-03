package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}


// 函数也是一种数据类型, 通过 type 给函数类型起名
type FuncType func(int, int) int

func main() {
	var result int
	result = Add(1,1)
	fmt.Println(result)

	var fTest FuncType
	fTest = Add
	result = fTest(1,1)
	fmt.Println(result)
}
