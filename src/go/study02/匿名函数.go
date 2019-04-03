package main

import "fmt"

func main() {
	a := 10
	str := "mike"

	f1 := func(){
		fmt.Println("a = ",a)
		fmt.Println("str = ",str)
	}

	f1()

	type FuncTtype func()

	var f2 FuncTtype
	f2 = f1
	f2()

	// 定义匿名函数同时调用
	func(){
		fmt.Printf("定义匿名函数同时调用1, a = %d,str = %s\n",a,str)
	}() // 后面对的 （） 表示调用该函数

	func(g,m int){
		fmt.Printf("定义匿名函数同时调用2, g = %d,m = %d\n",g,m)
	}(10, 20)
}
