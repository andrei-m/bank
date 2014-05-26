package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrei-m/bank/lib"
	"net/http"
	"strconv"
)

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

// Handle transaction listing
func handleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "[{\"Id\":0,\"Amount\":99,\"Date\":\"2014-05-23T00:00:00Z\"}]")
	} else {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Unsupported method", 405)
	}
}

func main() {
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/transaction/", handleTransaction)
	http.HandleFunc("/transactions", handleTransactions)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":1337", nil)
}
