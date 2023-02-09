package main

import (
	"log"

	"github.com/mchu1966/cli/cmd"
)

var name string

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
