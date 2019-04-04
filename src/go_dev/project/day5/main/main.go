package main

import (
	"../node"
	"fmt"
	"strconv"
)

func main() {

	nodeSplit := make([]*node.Node, 3)

	for i := 0; i < 3; i++ {
		n := node.NewInstance(strconv.Itoa(i), nil, i)
		nodeSplit[i] = n
		//nodeSplit = append(nodeSplit, n)
	}
	for i := 0; i < len(nodeSplit)-1; i++ {
		_, err := nodeSplit[i].SetNextNode(nodeSplit[i+1])
		if err != nil {
			return
		}
	}

	for _, no := range nodeSplit {
		fmt.Println(no.String())
	}

}
