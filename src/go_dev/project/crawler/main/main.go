package main

import (
	"../engine"
	"../lagou/parser"
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
func main01() {
	engine.Run(engine.Request{
		Url:       "https://www.lagou.com/gongsi/allCity.html?option=0-0-0-0",
		PaserFunc: parser.ParserCityList,
	})
}
func main() {
	resp, err := http.Get("https://www.lagou.com/gongsi/270-0-0")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", resp.StatusCode)
		return
	}
	// 自动转化编码
	encode := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	printCityList(all)

	//fmt.Printf("%s\n", all)
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(https://www.lagou.com/gongsi/[0-9]+-[0-9]+-[0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s,URL: %s\n", m[2], m[1])
	}

}
