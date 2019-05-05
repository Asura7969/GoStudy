package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetcher(url string) ([]byte, error) {
	time.Sleep(time.Second * 5)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// 自动转化编码
	fmt.Println(url)
	encode := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error :%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
