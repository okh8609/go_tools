package main

import (
	"log"

	"github.com/okh8609/go_tools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
