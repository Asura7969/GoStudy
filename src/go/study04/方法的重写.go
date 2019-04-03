package main

import "fmt"

type Peer01 struct {
	name string
	id int
}

func (p *Peer01)PrintValue()  {
	fmt.Printf("Peer01方法： name=%s,id=%d\n",p.name,p.id)
}

type Stty01 struct {
	Peer01
	address string
}

func (s *Stty01)PrintValue()  {
	fmt.Printf("Stty01方法：name=%s,id=%d\n",s.name,s.id)
}

func main() {
	s := Stty01{Peer01:Peer01{name:"test",id:1},address:"sh"}
	s.PrintValue()
	// 显示调用Peer01
	s.Peer01.PrintValue()
}
