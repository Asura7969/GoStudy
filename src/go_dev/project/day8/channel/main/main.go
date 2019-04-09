package main

import "fmt"

func main() {
	var chanMap chan map[string]string
	chanMap = make(chan map[string]string, 10)
	m := make(map[string]string, 16)

	m["1"] = "001"
	m["2"] = "002"

	chanMap <- m

	res := <-chanMap
	fmt.Println(res)

}
