package node

import (
	"fmt"
	"strconv"
)

type Node struct {
	content string
	next    *Node
	index   int
}

func (n *Node) String() string {
	str := "content:" + n.content + " index:" + strconv.Itoa(n.index)
	if n.next != nil {
		str += " next_index:" + strconv.Itoa(n.next.index)
	}
	return "{" + str + "}"
}

func NewInstance(content string, n *Node, i int) *Node {
	return &Node{content, n, i}
}

func (n *Node) GetNextNode() *Node {
	return n.next
}

func (n *Node) GetContent() string {
	return n.content
}

func (n *Node) GetIndex() int {
	return n.index
}

func (n *Node) SetNextNode(other *Node) (nt *Node, err error) {

	if n.next != nil {
		err = fmt.Errorf("this node has child! %s", n.next.content)
		return
	}
	n.next = other
	nt = n
	return
}

func SetIndex(n *Node, value int) {
	n.index = value
}

func AddNodes(n *Node, nodes []*Node) {
	currentIndex := n.index
	incrIndex := 1
	p := n

	if n.next != nil {
		nodes = MoveNodes(n.next, nodes)
		n.next = nil
	}

	for _, no := range nodes {
		currentIndex = currentIndex + incrIndex
		incrIndex += 1
		SetIndex(no, currentIndex)
		p.next = no
		p = no
	}
}

func MoveNodes(n *Node, moveNodes []*Node) []*Node {
	if n.next != nil {
		moveNodes = append(moveNodes, n)
		return MoveNodes(n.next, moveNodes)
	} else {
		moveNodes = append(moveNodes, n)
		return moveNodes
	}
}
