package engine

import (
	"../fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			log.Printf("Fetcher :err fetching url %s: %v", r.Url, err)
			continue
		}

		paserResult := r.PaserFunc(body)
		requests = append(requests, paserResult.Requests...)

		for _, item := range paserResult.Items {
			log.Printf("Got Item %v", item)

		}
	}
}
