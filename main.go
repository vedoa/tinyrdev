package main

import (
	"log"
	"os"

	"github.com/vedoa/tinyrdev/cli"
)

var v string

func main() {

	args := os.Args[1:]

	// If no arguments are passed, show help
	if len(args) == 0 {
		args = []string{}
	}

	err := cli.Execute(v, args)
	if err != nil {
		log.Fatal(err)
	}
}
