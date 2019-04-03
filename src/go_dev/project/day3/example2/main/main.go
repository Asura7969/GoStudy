package main

import "fmt"

func isNumber(n int) bool {

	var i,j,k int
	i = n % 10
	j = (n % 10) % 10
	k = (n % 100) % 10

	sum := i*i*i + j*j*j + k*k*k
	return sum == n

}
// 水仙花数
func main() {
	var n int
	fmt.Scanf("%d",&n)

	if isNumber(n) {
		fmt.Println(n,"is 水仙花")
	} else {
		fmt.Println(n,"is not 水仙花")
	}
}