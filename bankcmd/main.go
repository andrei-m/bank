package main

import (
	"fmt"
	"github.com/andrei-m/bank"
	"log"
	"net/http"
)

const (
	port = 1337
)

func main() {
	bank.SetupHandlers()
	log.Printf("listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
