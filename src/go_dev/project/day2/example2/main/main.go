package main

import (
	"fmt"
)
import "../add"

func main01() {
	fmt.Println("Name=", add.Name)
	fmt.Println("Age=", add.Age)
}

func main() {
	//a,_ := strconv.ParseInt("1100100",2,32)
	//b,_ := strconv.ParseInt("0144",8,32)
	//c,_ := strconv.ParseInt("64",16,32)
	//
	//println(a,b,c)
	//
	//println("0b" + strconv.FormatInt(a,2))
	//println("0" + strconv.FormatInt(a,8))
	//println("0x" + strconv.FormatInt(a,16))

	s := make([]int, 2, 10)
	for i, ss := range s {
		println(i, ss)
	}
	s = append(s, 100)
	println(s)

}
