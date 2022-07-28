package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jbrissier/gobring/api"
	"github.com/joho/godotenv"
)

func startSvelt() {

	npm := os.Getenv("NPM_PATH")

	if npm == "" {
		npm = "/usr/local/bin/npm"
	}

	cmd := exec.Command(npm, "run", "dev")
	cmd.Dir = "./svelte-app"
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	cmd.Wait()

}

func main() {

	// read sql
	godotenv.Load(".env")

	go startSvelt()
	api.StarServer()
}
