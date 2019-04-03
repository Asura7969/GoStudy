package main

//包：标识
//命名空间
//导入fmt包
import "fmt"

func main() {
	var i int = 10 //初始化
	var a, b int = 12, 13
	num := 20 //自动推导类型
	fmt.Printf("a=%d,b=%d,i=%d,num=%d", a, b, i, num)

	/*
	   fmt包作用：包含输入与输出。
	*/
	fmt.Println("传智播客")
    fmt.Print("hello world")
}
