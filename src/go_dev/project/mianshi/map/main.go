package main

import "sync"

// https://www.cnblogs.com/ghj1976/archive/2013/04/27/3047528.html
func main() {
	// 可能会出现fatal error: concurrent map read and map write.
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// 错误
// 问题出现在 Get方法
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

// 正确
//func (ua *UserAges) Get(name string) int {
//	ua.Lock()
//	defer ua.Unlock()
//	if age, ok := ua.ages[name]; ok {
//		return age
//	}
//	return -1
//}
