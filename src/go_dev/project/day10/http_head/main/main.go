package main

import (
	"fmt"
	"net"
	"net/http"
	time2 "time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {

	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (conn net.Conn, e error) {
					time := time2.Second * 2
					return net.DialTimeout(network, addr, time)
				},
			},
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head succ, status:%v\n", resp.Status)
	}
}
