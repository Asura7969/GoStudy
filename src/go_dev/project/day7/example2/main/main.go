package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 从终端读
func main01() {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed,err:", err)
		return
	}

	fmt.Printf("read str succ,ret: %s\n", readString)

}

// 从文件读
func main02() {
	file, err := os.Open("/Users/gongwenzhou/GolandProjects/GoStudy/README.md")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed,err:", err)
		return
	}

	fmt.Printf("read str succ,ret: %s\n", readString)
}

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

// 从文件中读取并计算字符
func main() {
	file, err := os.Open("/Users/gongwenzhou/GolandProjects/GoStudy/README.md")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	defer file.Close()

	var count CharCount

	for {
		reader := bufio.NewReader(file)
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("read string failed,err:", err)
			break
		}

		runeArr := []rune(str)

		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v == '0' && v == '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)

}
