package main

import (
	"fmt"
	"runtime"
	"sync"
)

// select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。
// 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
//
//1、select 中只要有一个case能return，则立刻执行。
//2、当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
//3、如果没有一个case能return则可以执行”default”块。
func main01() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}

}

/*
下面的迭代会有什么问题?
考点：chan缓存池


func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
*/
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// https://blog.csdn.net/samete/article/details/52751227

	ch := make(chan interface{}) // 解除注释看看！	无缓冲:除非有人(协程)一直从里面取数据,否则阻塞
	//ch := make(chan interface{},len(set.s)) //	有缓冲:只有等 chan 满了才会阻塞
	go func() {
		//time.Sleep(time.Second)
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter :", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2", "3"},
	}
	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
}
