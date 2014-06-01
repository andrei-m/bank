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
			http.Error(w, "Bad request body", 400)
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
		transactions := bank.LoadTransactions()
		j, err := json.Marshal(transactions)
		if err != nil {
			http.Error(w, "Failed to JSONify Transactions", 500)
		} else {
			fmt.Fprintf(w, string(j))
		}
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
