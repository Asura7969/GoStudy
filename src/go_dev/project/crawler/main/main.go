package main

import (
	"../engine"
	"../lagou/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:       "https://www.lagou.com/gongsi/allCity.html?option=0-0-0-0",
		PaserFunc: parser.ParserCityList,
	})

}
