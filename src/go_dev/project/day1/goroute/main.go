package main

import(
	"fmt"
	"time"
)

func main01() {

	//for i := 0; i < 100; i++ {
	//	go test_goroute(i)
	//}

	c := make(chan int,3)
	c <- 1
	fmt.Println("chan:", <- c)
	// fmt.Println("chan:", <- c)  报错

	time.Sleep(time.Second)
}
func write(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func read(c chan int) {
	for {
		fmt.Println("go:",<-c)
		time.Sleep(1)
	}
}
func readgo(c chan int)  {
	for v := range c{
		fmt.Println("并发通信:",v)
	}
}

func main() {

	//c := make(chan int)
	//go write(c)
	//go read(c)
	//go readgo(c)
	//time.Sleep(2)


	a := make(chan int, 15)
	for i := 0; i < 10; i++ {
		a <- i
	}

	// 不关闭会报错
	close(a)

	for {
		if v, ok := <-a; !ok {
			fmt.Println("不存在")
			break
		} else {
			fmt.Println(v)
			//fmt.Println(ok)
		}
	}
}