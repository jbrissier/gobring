package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbrissier/gobring/server"
)

func main() {

	go func() {
		cmd := exec.Command("/usr/local/bin/npm", "run", "dev")
		cmd.Dir = "./svelte-app"
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
		cmd.Wait()
	}()
	server.StarServer()
}
