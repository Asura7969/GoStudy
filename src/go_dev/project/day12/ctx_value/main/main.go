package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 0
	}

	fmt.Printf("ret:%d\n", ret)

	s, ok := ctx.Value("session").(string)
	if !ok {
		s = "default"
	}
	fmt.Println(s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 1231221)
	ctx = context.WithValue(ctx, "session", "sdk")
	process(ctx)
}
