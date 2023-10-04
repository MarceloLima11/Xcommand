package main

import (
	"log"

	"github.com/MarceloLima11/show-files/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
