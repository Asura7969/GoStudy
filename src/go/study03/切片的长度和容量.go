package main

import "fmt"

// 切片的创建
func main() {
	// 自动推导类型，同时初始化
	s1 := []int{1,2,3}
	fmt.Println("s1 = ",s1)

	// 借助make函数，格式 make (切片类型，长度，容量)
	s2 := make([]int,5,10)
	fmt.Printf("len = %d,cap = %d\n",len(s2),cap(s2))

	// 没有指定容量，容量和长度一样
	s3 := make([]int,5)
	fmt.Printf("len = %d,cap = %d\n",len(s3),cap(s3))

}


func main01() {

	a := [5]int{0,1,2,3,4}
	s := a[0:3:5]

	fmt.Println("s = ",s)
	fmt.Println("len(s) = ",len(s)) // 长度 3 - 0
	fmt.Println("cap(s) = ",cap(s)) // 容量 5 - 0

	s = a[1:2:3]
	fmt.Println("s = ",s)
	fmt.Println("len(s) = ",len(s)) // 长度 2 - 1
	fmt.Println("cap(s) = ",cap(s)) // 容量 3 - 1


	// 切片和数组的区别
	// 切片 "[]" 里面可以为空或者 "..."
	h := []int{}
	fmt.Printf("切片h => len(h):%d,cap(h):%d\n",len(h),cap(h))

	h = append(h,11)
	fmt.Printf("切片h append => len(h):%d,cap(h):%d",len(h),cap(h))


}
