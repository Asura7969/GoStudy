package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func main() {
	//wait01()

	// 等待多个任务结束,建议使用此方法
	//waitMutiJob()

	// 可在多处使用Wait阻塞,他们都能接收到通知
	//wait02()

	// 计算 goroutine 执行时间
	//n := runtime.GOMAXPROCS(0)
	//fmt.Println(n)
	//test(n)
	//test2(n)

	// Local Storage,goroutine 不支持优先级设置、无法获取编号、没有局部存储(TLS),甚至连返回值都会抛弃
	//local_stroage()

	// 暂停,释放线程取执行其他任务,当前任务被放回队列,等待下次调度时恢复执行
	//goSched()

	// 立即终止当前任务,运行时确保所有已注册延迟调用被执行
	goExit()

}

func goExit() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer fmt.Println("a")

		func() {
			defer func() {
				fmt.Println("b", recover() == nil) // 执行,recover 返回 nil
			}()

			func() {
				fmt.Println("c")
				runtime.Goexit()
				fmt.Println("c done.") // 不会执行
			}()

			fmt.Println("b done") // 不会执行
		}()

		fmt.Println("a done") // 不会执行
	}()

	<-exit
	fmt.Println("main exit")
}

func goSched() {
	runtime.GOMAXPROCS(1) // 设置执行使用的核心数
	exit := make(chan struct{})

	go func() { // 任务 a
		defer close(exit)

		go func() { // 任务b 放在此处是为了任务a优先执行
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)
			if i == 1 { // 让出当前线程,调度执行b
				runtime.Gosched()
			}
		}
	}()

	<-exit
}

func local_stroage() {
	var (
		wg sync.WaitGroup
		gs [5]struct { // 用于实现类似TLS功能
			id     int // 编号
			result int // 返回值
		}
	)

	for i := 0; i < len(gs); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", gs)

}

// 测试目标函数
func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += 1
	}
	println(x)
}

// 循环执行
func test(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

// 并发执行
func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}

	wg.Wait()
}

func wait02() {
	var (
		wg sync.WaitGroup
	)

	wg.Add(1)

	go func() {
		wg.Wait() // 等待归零,解除阻塞
		fmt.Println("wait exit...")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("done.")
		wg.Done()
	}()

	//time.Sleep( 2 * time.Second)
	wg.Wait() // 等待归零,解除阻塞
	fmt.Println("main exit")
}

func waitMutiJob() {
	var (
		wg sync.WaitGroup
	)

	for i := 0; i < 5; i++ {
		wg.Add(1) // 累加计数
		go func(id int) {
			//wg.Add(1) 不建议在此处设置,来不及设置
			defer wg.Done() // 递减计数
			time.Sleep(2 * time.Second)
			println("goroutine", id, "done.")
		}(i)
	}
	fmt.Println("main ...")
	wg.Wait() // 阻塞,直到计数归零
	fmt.Println("main exit ...")
}

func wait01() {
	var (
		exit chan struct{}
	)

	exit = make(chan struct{}, 100)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine doing...")
		// 关闭chain,发出信号
		close(exit)
	}()

	fmt.Println("main doing...")
	<-exit // 接受通道关闭,立即解除阻塞
	fmt.Println("exit main")
}
