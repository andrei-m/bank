package bank

import (
	"encoding/json"
	"fmt"
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

		transaction := LoadTransaction(id)
		fmt.Fprintf(w, transaction.JSON())
	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var transaction Transaction
		err := decoder.Decode(&transaction)
		if err != nil {
			http.Error(w, "Bad request body", 400)
			return
		}

		err = transaction.Save()
		if err != nil {
			//TODO: a missing date is a 4xx, the other errors are probably 5xx
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprintf(w, transaction.JSON())
	} else if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Path[len("/transaction/"):])
		if err != nil {
			http.Error(w, "Bad or missing transaction id in route", 400)
			return
		}

		transaction := LoadTransaction(id)
		err = transaction.Delete()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete transaction %d", id), 500)
			return
		}

		fmt.Fprintf(w, "Deleted %d", id)
	} else {
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "Unsupported method", 405)
	}
}

// Handle transaction listing
func handleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		transactions := LoadTransactions()
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

func SetupHandlers() {
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/transaction/", handleTransaction)
	http.HandleFunc("/transactions", handleTransactions)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
}
