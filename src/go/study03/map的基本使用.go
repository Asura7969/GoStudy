package main

import "fmt"

func main() {
	// 定义一个变量, 类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ",m1)
	// 对于map只有len,没有cap
	fmt.Println("len = ",len(m1))

	// 可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ",m2)
	fmt.Println("len = ",len(m2))

	// 可以指定长度,只是指定容量，但是里面却是一个数据也没有
	m3 := make(map[int]string)
	m3[1] = "mike"
	m3[2] = "java"
	m3[3] = "go"
	fmt.Println("m3 = ",m3)
	fmt.Println("len = ",len(m3))

	// 初始化
	m4 := map[int]string{1:"mike",2:"go",3:"java"}
	fmt.Println("m4 = ",m4)
	delete(m4,3)
	fmt.Println("删除操作后 m4 = ",m4)

}
