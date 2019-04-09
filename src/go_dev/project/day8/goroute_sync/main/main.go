package main

import "fmt"

// 变量名 <- chan int : 只读
// 变量名 chan <- int : 只写
func calc(taskChan <-chan int, ranChan chan int, exitChan chan<- bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			ranChan <- v
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)

	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 1000; i++ {
			intChan <- i
		}

		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	// 等待所有计算的goroute全部退出
	for i := 0; i < 8; i++ {
		<-exitChan
		fmt.Println("wait goroute", i)
	}

	close(resultChan)

	for _ = range resultChan {
		//fmt.Println(r)
	}

}
