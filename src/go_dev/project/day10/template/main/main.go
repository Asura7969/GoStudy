package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	age  string
}

func main() {
	t, err := template.ParseFiles("F:/golangWorkspace/GoStudy/src/go_dev/project/day10/template/main/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
