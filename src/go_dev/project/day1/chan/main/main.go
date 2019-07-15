package main

import (
	"fmt"
	"time"
)

func add(a int, b int, p chan int) {
	sum := a + b
	time.Sleep(3 * time.Second)
	p <- sum
}

func main() {
	var pip chan int

	pip = make(chan int, 1)
	go add(1, 2, pip)

	summ := <-pip
	fmt.Println("sum = ", summ)

	// pipe <- 3
	// var t1 int
	// t1 = <- pipe
	// fmt.Println(t1)

	fmt.Println(len(pip))

}
