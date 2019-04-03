package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	func(){
		// 闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("inner, a = %d,str = %s\n",a,str)
	}()
	fmt.Printf("outer, a = %d,str = %s\n",a,str)
}
