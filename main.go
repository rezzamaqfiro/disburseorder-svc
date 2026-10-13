package main

import (
	"log"

	"disburseorder-svc/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
