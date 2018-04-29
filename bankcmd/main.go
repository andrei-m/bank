package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrei-m/bank"
	"github.com/phogolabs/parcello"

	_ "github.com/andrei-m/bank/static"
)

const port = 1337

func main() {
	bank.SetupHandlers(parcello.Root("/"))
	log.Printf("listening on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
