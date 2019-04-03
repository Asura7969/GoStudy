package main

import "fmt"

func main() {
	//num := 4

	// 支持一个初始化语句
	switch num := 4; num {
	case 1:
		fmt.Println("1")
		//break // 跳出switch , 默认包含
		fallthrough //强制执行后边语句（无条件执行）,只执行一个
	case 2:
		fmt.Println("2")
		//fallthrough
	case 3 ,10, 100: // 可以写多个值
		fmt.Println("3 ,10, 100")
		//fallthrough
	case 4:
		fmt.Println("4")
		fallthrough
	default:
		fmt.Println("......")
	}

	score := 85
	switch  {
	case score >= 90: // case 后面可以放条件
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("一般")
	case score >= 70:
		fmt.Println("差")
	default:
		fmt.Println(".....")

	}
}
