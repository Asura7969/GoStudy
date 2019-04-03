package main

import "fmt"

func main() {
	// 给 int64 起别名为 bigint
	type bigint int64

	var a bigint

	fmt.Printf("a type is %T\n", a)

	type(
		long int64
		char byte
	)
	var b long = 11
	var ch char = 'a'

	fmt.Printf("b = %d, ch = %c\n",b,ch)
}
