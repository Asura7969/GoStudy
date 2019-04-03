package main

import "fmt"

type Su struct {
	name string
	id int
}

func main() {
	i := make([]interface{},3)
	i[0] = 1
	i[1] = "str"
	i[2] = Su{"su",1}
	fmt.Println("if...")
	for index,data := range i {
		// 第一个返回的是指，第二个返回的是结果的真假
		if value,ok := data.(int);ok == true{
			fmt.Printf("x[%d]类型为int，内容为%d\n",index,value)
		} else if value,ok := data.(string);ok == true{
			fmt.Printf("x[%d]类型为string，内容为%d\n",index,value)
		} else if value,ok := data.(Su);ok == true{
			fmt.Printf("x[%d]类型为Su，内容为%d\n",index,value)
		}
	}
	fmt.Println("switch...")
	for index,data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int，内容为%d\n",index,value)
		case string:
			fmt.Printf("x[%d]类型为string，内容为%d\n",index,value)
		case Su:
			fmt.Printf("x[%d]类型为Su，内容为%d\n",index,value)
		}
	}
}
