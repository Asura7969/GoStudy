package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//notice()

	// 异步
	//asynNotice()

	// 判断通道是同步还是异步;通过内置函数 len,cap 返回缓冲区大小和当前已缓冲数量;而对于同步通道则都返回0
	//judgeChan()

	// 使用 ok-idom 或 range 模式处理数据
	//forChan()
	//rangeChan()

	// 接受指定事件后关闭chan
	//specialChan()

	// 单向
	//oneWayChan()

	// select 使用
	//selectChan()

	// 等待所有通道消息处理结束（close）,可将已完成的通道设置成nil，这样就不会被阻塞,不再被 select 选中
	//closeAllChan()

	// 当所有通道不可用时,select会执行default语句
	//defaultCase()

	// 使用工厂方法将 goroutine 和通道绑定
	//factoryCase()

	// 用通道实现信号量
	semaphoreForChan()

}

// ======================== chan 可用作ID Generator 或 pool ========================

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p: // 返回
	default: // 返回失败,新建
		v = make([]byte, 10)
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b: // 放回
	default: // 放回失败
	}
}

// ========================================================================
func factoryCase() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data) // 通道关闭,发出结束通知
	r.Wait()      // 等待接受者处理结束
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)

	go func() {
		defer r.Done()
		for x := range r.data { // 接受消息,直到通道被关闭
			fmt.Println("recv:", x)
		}
	}()

	return r
}

type receiver struct {
	sync.WaitGroup
	data chan int
}

func defaultCase() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for true {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data", x)
			default: // 避免 select 阻塞
				fmt.Println("default")
			}

			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)

	c <- 100
	close(c)

	<-done

}

func closeAllChan() {
	var wg sync.WaitGroup

	wg.Add(3)

	a, b := make(chan int), make(chan int)
	go func() { // 接收端
		defer wg.Done()

		for {
			select {
			case x, ok := <-a:
				if !ok { // 如果通道关闭,则设置为 nil
					a = nil
					break
				}
				fmt.Println("a", x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				fmt.Println("b", x)
			}

			if a == nil && b == nil { // 全部结束,退出循环
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()

}

func selectChan() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() { // 接收端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)

			select { // 随机选择可用的 channel 接受数据
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok { // 如果任意通道关闭,则终止接受
				return
			}

			fmt.Println(name, x) // 输出接受的数据信息
		}
	}()

	go func() { // 发送端
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { // 随机选择发送channel
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	wg.Wait()
}

func oneWayChan() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()

		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	wg.Wait()

	// close(recv)				close不能作用于接收端
}

func specialChan() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println(id, ":ready.")
			<-ready
			fmt.Println(id, ":running...")
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Ready? Go!")

	close(ready)

	wg.Wait()

}

func rangeChan() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)

	<-done
}

func forChan() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			x, ok := <-c
			if !ok { // 据此判断通道是否被关闭
				return
			}
			fmt.Println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c)

	<-done
}

func judgeChan() {
	a, b := make(chan int), make(chan int, 3)

	b <- 1
	b <- 2

	fmt.Println("a(同步通道):", len(a), cap(a))
	fmt.Println("b(异步通道):", len(b), cap(b))
}

func asynNotice() {
	c := make(chan int, 3) // 创建3个缓冲槽的异步通道

	c <- 3 // 缓冲区未满,不会阻塞
	c <- 2

	fmt.Println(<-c) // 缓冲区尚有数据,不会阻塞
	fmt.Println(<-c)

	// todo:多数时候,异步通道有助于提升性能,减少排队阻塞
}

func notice() {
	done := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c
		fmt.Println(s)
		close(done)
	}()

	c <- "hi!"
	<-done
}
