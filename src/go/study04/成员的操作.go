package main

import "fmt"

type Person01 struct{
	name string
	sex byte
	age int
}

type Student01 struct{
	Person01
	id int
	address string
}

func main() {
	s1 := Student01{Person01{"ls",'m',20},1,"shanghai"}
	fmt.Println(s1.name,s1.age,s1.Person01)

	s1.age = 100
	fmt.Println("s1 = ",s1)
	s1.address = "bj"
	s1.Person01 = Person01{"zs",'w',200}
	fmt.Println("s1 = ",s1)


}
