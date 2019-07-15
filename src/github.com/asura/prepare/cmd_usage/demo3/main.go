package mian

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {

	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		output     []byte
		err        error

		resultChain chan *result
		res         *result
	)

	resultChain = make(chan *result, 1000)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		cmd = exec.Command("go", "env")

		output, err = cmd.CombinedOutput()

		resultChain <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)

	cancelFunc()

	res = <-resultChain

	fmt.Println(res.err, string(res.output))
}
