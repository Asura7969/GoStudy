package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		url  string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)

	fmt.Println(urlProcess(url))
	fmt.Println(pathProcess(path))

}

func urlProcess(url string) string {
	if !strings.HasPrefix(url, "http://") {
		return "http://" + url
	}
	return url
}

func pathProcess(path string) string {
	if !strings.HasSuffix(path, "/") {
		return path + "/"
	}
	return path
}
