package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i ++ {
		s[i] = rand.Intn(100)
	}
}

func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n - 1; i ++ {
		for j := 0; j < n - 1 - i; j ++ {
			if s[j] > s[j+1] {
				s[j] , s[j+1] = s[j+1] , s[j]
			}
		}
	}
}

// 切片作为参数是引用传递,里面改了,会影响外面的实参
func main() {

	v := 10
	// 创建一个切片，len为v
	s := make([]int,v)
	// 初始化
	InitData(s)
	fmt.Println("排序前：",s)
	// 冒泡排序
	BubbleSort(s)
	fmt.Println("排序后：",s)
	
}
