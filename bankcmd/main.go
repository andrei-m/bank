package main

import (
	"fmt"
	"github.com/andrei-m/bank"
	"net/http"
)

func main() {
	bank.SetupHandlers()
	err := http.ListenAndServe(":1337", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
