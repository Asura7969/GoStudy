package main

import (
	"fmt"
	"regexp"
)

const text string = `
my email is 1402367@qq.com
asdasd@gmail.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match:= re.FindString(text)
	submatch := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)

	fmt.Println(submatch)

}
