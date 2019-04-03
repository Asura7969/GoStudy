package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for  {// for 后面不写，默认一直执行
		i ++
		time.Sleep(time.Second)
		if i == 5 {
			break //跳出循环
			//continue // 跳出本次循环
		}

		fmt.Println(i)

	}
}
