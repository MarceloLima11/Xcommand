package main

import (
	"log"

	"github.com/MarceloLima11/Xcommand/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
