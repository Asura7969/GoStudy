package main

import "fmt"

func swapAddress(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {
	a,b := 1,100
	swapAddress(&a,&b)

	fmt.Printf("main  =>  a:%d, b:%d\n",a,b)
}
