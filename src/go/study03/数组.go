package main

import "fmt"

func main() {

	// 定义一个数组,[10]int 和 [5]int是不同类型
	// "[]"里面的数组为数组的长度
	var a [10]int
	var b [5]int

	fmt.Printf("len(a):%d,len(b):%d",len(a),len(b))

	//n:=10
	//var c [n]int  // err,"[]"里面必须是常量

	// 赋值每个元素
	for i := 0;i < len(a); i++ {
		a[i] = i + 1
	}

	// 第一个返回下标,第二个返回元素
	for i,data := range a{
		fmt.Printf("a[%d] = %d\n",i,data)
	}

	// 数组初始化
	var c [5]int = [5]int{1,2,3,4,5}
	// c := [5]int{1,2,3,4,5}
	fmt.Println("c = ",c)
	// 部分初始化
	c = [5]int{1,2,3}
	fmt.Println("c = ",c)
	// 指定元素初始化
	c = [5]int{1:1,2:2}
	fmt.Println("c = ",c)

	// 数组比较
	d := [5]int{1,2,3,4,5}
	e := [5]int{1,2,3,4,5}
	f := [5]int{1,2,3}
	fmt.Println("d == e", d == e)
	fmt.Println("d == f", d == f)

}
