package main

import "fmt"

//定义函数：形参
func Test(num1, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

func main() {
	s := Test(3, 5) //实参
	fmt.Println(s)

	/*
		//乘法口诀表
		for i := 1; i <= 9; i++ {
			for j := 1; j <= 9; j++ {
				fmt.Printf("%d*%d=%d\t", i, j, i*j)
			}
			fmt.Println("")
		}*/

	/*

			要求用户输入用户名和密码，只要不是admin、888888就一直提示用户名,
		             密码错误,请重新输入

	*/
	/*
		var userName string
		var userPwd string

		for {
			fmt.Println("请输入用户名:")
			fmt.Scan(&userName)
			fmt.Println("请输入密码:")
			fmt.Scan(&userPwd)
			if userName == "admin" && userPwd == "888888" {
				fmt.Println("用户名密码输入正确")
				break
			} else {
				fmt.Println("用户名密码错误,请重新输入")
			}
		}*/

	/*
		var score int = 80
		switch score {
		case 90:
			fmt.Println("A")
		case 80:
			fmt.Println("B")
			fallthrough
		case 70:
			fmt.Println("C")

		default:
			fmt.Println("D")
		}*/

	/*

	   成绩>=90 ：A
	   90>成绩>=80 ：B
	   80>成绩>=70 ：C
	   70>成绩>=60 ：D
	   成绩<60  ：E

	*/
	/*
		var score float64
		fmt.Println("请输入成绩:")
		fmt.Scan(&score)
		if score >= 90 {
			fmt.Println("A")
		} else if score >= 80 {
			fmt.Println("B")
		} else if score >= 70 {
			fmt.Println("C")
		} else if score >= 60 {
			fmt.Println("D")
		} else {
			fmt.Println("E")
		}*/

	/*

		var chinese int = 90
		var english int = 80
		var math int = 69
		fmt.Printf("sum=%d,avg=%.2f", chinese+english+math, float64(chinese+math+english)/3)*/
	/*
		a, b := 10, 20
		a, b = b, a //注意：这里没有冒号
		fmt.Printf("a=%d,b=%d", a, b)*/
}
