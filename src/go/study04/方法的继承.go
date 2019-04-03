package main

import "fmt"

type Peer struct {
	name string
	id int
}

func (p *Peer)PrintValue()  {
	fmt.Printf("name=%s,id=%d\n",p.name,p.id)
}

type Stty struct {
	Peer
	address string
}

func main() {
	s := Stty{Peer:Peer{name:"test",id:1},address:"sh"}
	s.PrintValue()
}
