package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)
	if expr, err = cronexpr.Parse("* * * * * * *"); nil != err {
		fmt.Println(err)
		return
	}
	// 当前时间
	now = time.Now()
	// 下次调度时间
	nextTime = expr.Next(now)

	//等待这个定时器超时 下次调度时间 减 当前时间(now) = 等待时间间隔
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println(nextTime, ":被调度了")
	})

	time.Sleep(10 * time.Second)
}
