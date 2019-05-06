package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe = `<a href="(https://www.lagou.com/gongsi/[0-9]+-[0-9]+-[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.PaserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.PaserResult{}
	for _, m := range matches {

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			PaserFunc: engine.NilParser,
		})
	}
	return result
}
