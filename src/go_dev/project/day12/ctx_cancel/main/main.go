package main

import (
	"context"
	"fmt"
	"time"
)

/**
ctx 控制 goroute退出
*/
func test() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for true {
				select {
				case <-ctx.Done():
					fmt.Println("end")
					return
				case dst <- n:
					n++
				}

			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func main() {
	test()
	time.Sleep(time.Hour)
}
