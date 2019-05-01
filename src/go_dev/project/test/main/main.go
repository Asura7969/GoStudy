package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)

}

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

/**
字符串修改
*/
func modifyString() {
	// 修改字符串,需将其转换为可变类型([]byte 或 []rune),待完成后再转换回来（都需要重新分配内存，并复制数据）
	s := "hello,world!"
	pp("s: %x\n", &s)

	bs := []byte(s)
	s2 := string(bs)

	pp("string to []byte, bs:%x\n", &bs)
	pp("[]byte to string, s2:%x\n", &s2)

	rs := []rune(s)
	s3 := string(rs)

	pp("string to []rune, rs:%x\n", &rs)
	pp("[]rune to string, s3:%x\n", &s3)

	// 使用非安全的方法(以上方法会降低性能)
	test_str := []byte("hello,world!")
	ss := toString(test_str)
	pp("bs: %x\n", &test_str)
	pp("s: %x\n", &ss)
}

type User struct {
	name string
	age  byte
}

/*
	数组初始化
		数组是值类型，赋值和传参都会拷贝整个数组数据，可以使用指针或切片，以此避免数据复制
*/
func initArray() {
	// 初始化数组，并指定索引位的初始值
	num1 := [2]int{1: 10}
	fmt.Println(num1)

	d := [...]User{
		{"Tom", 20},
		{"Mary", 21},
	}
	fmt.Printf("%#v\n", d)
}

// 切片拷贝
func copySilce() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[5:8]
	n := copy(s[4:], s1) // copy( [4,5,6,7,8,9], [5,6,7] )
	fmt.Println(n, s)

	s2 := make([]int, 6) // new 返回指针, make返回初始化之后类型的引用
	fmt.Println("s2:", s2)
	n = copy(s2, s)
	fmt.Println(n, s2)
}

func main() {

	// modifyString()
	//initArray()
	copySilce()

}
