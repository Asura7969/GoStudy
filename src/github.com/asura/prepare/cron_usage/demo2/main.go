package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CrontabJob struct {
	jobName  string
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		now           time.Time
		scheduleTable map[string]*CrontabJob
	)

	now = time.Now()
	jobOne := createCrontabJob("job1", "*/5 * * * * * *", now)
	jobTwo := createCrontabJob("job2", "*/10 * * * * * *", now)

	scheduleTable = make(map[string]*CrontabJob)

	scheduleTable["job1"] = jobOne
	scheduleTable["job2"] = jobTwo

	// 调度任务线程
	go func() {
		var (
			jobName string
			cronJob *CrontabJob
			now     time.Time
		)
		// 定时检查一下调度表
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Equal(now) || cronJob.nextTime.Before(now) {
					fmt.Println("-------------------------------------------")
					// 启动一个线程,执行这个任务
					go func(jobName string) {
						fmt.Println(jobName, "任务执行")
					}(jobName)

					// 计算下一次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, " 任务 的下次执行时间:", cronJob.nextTime)

					fmt.Println("-------------------------------------------")
				}
			}

			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: // 将在100毫秒后可读,返回
			}

			//time.Sleep(100 * time.Millisecond)
		}

	}()

	time.Sleep(100 * time.Second)

}

func createCrontabJob(jobName string, expression string, now time.Time) (resultCrobJob *CrontabJob) {
	expr := cronexpr.MustParse(expression)
	resultCrobJob = &CrontabJob{
		jobName:  jobName,
		expr:     expr,
		nextTime: expr.Next(now),
	}
	return
}
