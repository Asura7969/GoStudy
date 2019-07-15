package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	cmd = exec.Command("go", "env")

	if output, err = cmd.CombinedOutput(); err == nil {
		fmt.Println(string(output))
	} else {
		fmt.Println(err)
	}
}
