package main

import "fmt"

func main() {
	str := "aaa"

	for i := 0; i< len(str);i++ {
		fmt.Println(str[i])
	}

	for i,data := range str{
		fmt.Printf("str[%d] = %c \n",i,data)
	}

	for i := range str{ // 第2个返回值默认丢弃
		fmt.Printf("str[%d] = %c \n",i, str[i])
	}
}
