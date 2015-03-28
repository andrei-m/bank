package main

import (
	"github.com/andrei-m/bank"
	"log"
	"net/http"
)

func main() {
	bank.SetupHandlers()
	log.Println("listening on port 1337")
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		log.Fatal(err)
	}
}
