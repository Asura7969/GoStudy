package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64((n)))); i ++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
// 求素数
func main() {
	var n int
	var m int

	fmt.Scanf("%d%d%s", &n, &m)
	for i := 0; i < m; i ++ {
		if isPrime(i) {
			fmt.Printf("%d\n",i)
			continue
		}
	}


}
