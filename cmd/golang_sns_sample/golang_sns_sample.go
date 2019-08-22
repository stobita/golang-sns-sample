package main

import (
	"log"

	"github.com/stobita/golang-sns-sample/internal/server"
)

func main() {
	log.Fatal(server.Run())
}
