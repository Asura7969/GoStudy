package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}

type Student struct{
	Person
	id int
	address string
}

func main() {
	var s1 Student = Student{Person{"ls",'m',20},1,"shanghai"}
	fmt.Println("s1 = ",s1)

	// 自动推导类型
	s2 := Student{Person{"ls",'m',20},1,"shanghai"}
	fmt.Println("s2 = ",s2)

	// 指定成员初始化,没有初始化的成员为0
	s3 := Student{id:1}
	fmt.Println("s3 = %+v",s3)

	s4 := Student{Person:Person{name:"ls"},id:1}
	fmt.Println("s4 = %+v",s4)

}


