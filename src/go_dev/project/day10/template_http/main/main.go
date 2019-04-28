package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Person struct {
	Name string
	age  string
}

var myTemplate *template.Template

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	p := Person{Name: "Mary", age: "31"}
	// 输出到浏览器
	myTemplate.Execute(w, p)
	// 输出到终端
	myTemplate.Execute(os.Stdout, p)

	file, e := os.OpenFile("c:/test.log", os.O_CREATE|os.O_WRONLY, 0755)
	if e != nil {
		fmt.Println("open file err:", e)
	}
	// 输出到文件
	myTemplate.Execute(file, p)

}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file,err", err)
		return
	}
	return

}

func main() {
	initTemplate("F:/golangWorkspace/GoStudy/src/go_dev/project/day10/template_http/main/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
