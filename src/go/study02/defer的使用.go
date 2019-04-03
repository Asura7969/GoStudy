package main

import "fmt"

func ttest(x int) {
	result := 100 / x
	fmt.Println(result)
}

// defer 在函数结束之前调用
// 多个defer的调用顺序,先加载的最后执行（先进后出【输出】）
func main() {

	//defer fmt.Println("aaaaaaaaaaa")
	//defer fmt.Println("bbbbbbbbbbb")
	//defer ttest(0)
	//defer fmt.Println("ccccccccccc")

	// defer与匿名函数的使用
	deferTest()

}

func deferTest() {
	a:=10
	b:=20

	defer func(a,b int){
		fmt.Printf("里面的 ：a = %d,b = %d \n",a,b)
	}(a,b)

	a = 100
	b = 200

	fmt.Printf("外面的 ：a = %d,b = %d \n",a,b)
}