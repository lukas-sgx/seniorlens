package main

import (
	"log"

	"github.com/lukas-sgx/seniorlens/cmd/seniorlens/cli"
)

func main() {
	err := cli.Controller()

	if err != nil {
		log.Fatal(err)
	}
}