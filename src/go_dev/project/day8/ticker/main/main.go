package main

import (
	"fmt"
	"time"
)

func main() {

	d := time.Duration(time.Second * 2)
	ticker := time.NewTicker(d)
	for v := range ticker.C {
		fmt.Println("hello", v)

	}
	ticker.Stop()

	// 一秒之后执行
	select {
	case <-time.After(time.Microsecond):
		fmt.Println("执行")
	}
}
