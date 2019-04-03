package main

import "fmt"

// 注意，函数名大写表示public,小写表示private
func main() {

	//myFunc()
	//commonFunc(1)
	//fmt.Println(callBackFunc(1,2))
	//
	//argusFunc(1,2,3)

	// 参数传递
	//myFunc2(1,2,3,4)
	//fmt.Println(test2())


	fmt.Printf("%d \n", CommonSumFunc())
	fmt.Printf("%d", SpecalSumFunc(100))
}
// 无参无返回值
func myFunc() {
	fmt.Println("无参无返回函数")
}

// 普通参数列表
func commonFunc(a int) {
	fmt.Println(a)
}

func commonFunc2(a,b int,d string) {
	fmt.Println(a,b)
}

// 有参数有返回函数
func callBackFunc(a int, b int) bool {
	return a > b
}

// 常用写法,给 result 变量赋值
func test1()(result int){
	result = 666
	return
}

// 多个返回值
func test2()(a,b,c int){
	a,b,c = 1,2,3
	return
}

// 不定参数列表
func argusFunc(args ...int){
	fmt.Println(len(args))
}

func myFunc2(args ...int) {
	// 注意区别
	test(args[2:]...)
	test(args[:2]...)
}

func test(args ...int) {
	for _, data := range args {
		fmt.Println("data = ",data)
	}
}


func CommonSumFunc() (sum int) {
	for i := 1; i <= 100; i ++{
		sum += i
	}
	return
}

func SpecalSumFunc(a int) int {
	if a == 1 {
		return 1
	}
	return a + SpecalSumFunc(a - 1)
}
