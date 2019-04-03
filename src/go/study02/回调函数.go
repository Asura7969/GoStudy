package main

import "fmt"

type TwoArgsFuncType func(int, int) int

func Add01(a, b int) int {
	return a + b
}

func Mul01(a, b int) int {
	return a * b
}

func test09(a, b int, fTest TwoArgsFuncType) {
	fmt.Println(fTest(a,b))
}

func main() {

	test09(2,4, Add01)
	test09(2,4, Mul01)
}
