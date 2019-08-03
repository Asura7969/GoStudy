package main

import (
	"fmt"
)

// 以下代码能编译过去吗？为什么？

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//var peo People = Stduent{}
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
