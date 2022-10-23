package main

import (
	"log"

	"github.com/badcode256/example_go_hexagonal/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Start(); err != nil {
		log.Fatalf("Error start server : " + err.Error())
	}
}
