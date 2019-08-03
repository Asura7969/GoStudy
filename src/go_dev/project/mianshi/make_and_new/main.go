package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法		range是创建了每个元素的副本,而不是直接返回每个元素的引用
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	//for k,v:=range m{
	//	println(k,"=>",v.Name)
	//}
	fmt.Println("======正确=======")

	// 正确写法
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}

}

func main() {
	i := new(int) // new 返回地址
	fmt.Println(i)
	pase_student()

}
