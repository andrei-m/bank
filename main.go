package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrei-m/bank/lib"
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO")
}

// Handle CRUD for transactions
func handleTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Path[len("/transaction/"):])
		if err != nil {
			http.Error(w, "Bad or missing transaction id in route", 400)
			return
		}

		transaction := bank.LoadTransaction(id)
		fmt.Fprintf(w, transaction.JSON())
	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var transaction bank.Transaction
		err := decoder.Decode(&transaction)
		if err != nil {
			fmt.Println("Bad request body")
			return
		}

		transaction.Save()
		fmt.Fprintf(w, transaction.JSON())
	} else {
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Unsupported method", 405)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/transaction/", handleTransaction)
	http.ListenAndServe(":1337", nil)
}
