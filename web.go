package bank

import (
	"encoding/json"
	"fmt"
	"log"
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

		transaction := loadTransaction(id)
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(transaction); err != nil {
			log.Printf("Encode(%v): %v", transaction, err)
			http.Error(w, "failed to write transaction response", 500)
		}
	} else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var transaction transaction
		if err := decoder.Decode(&transaction); err != nil {
			http.Error(w, "Bad request body", 400)
			return
		}

		if err := transaction.Save(); err != nil {
			//TODO: a missing date is a 4xx, the other errors are probably 5xx
			http.Error(w, err.Error(), 400)
			return
		}
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(transaction); err != nil {
			log.Printf("Encode(%v): %v", transaction, err)
			http.Error(w, "failed to write transaction response", 500)
		}
	} else if r.Method == "DELETE" {
		id, err := strconv.Atoi(r.URL.Path[len("/transaction/"):])
		if err != nil {
			http.Error(w, "Bad or missing transaction id in route", 400)
			return
		}

		transaction := loadTransaction(id)

		if err := transaction.Delete(); err != nil {
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
		transactions := loadTransactions()
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(transactions); err != nil {
			log.Printf("Encode(%v): %v", transactions, err)
			http.Error(w, "failed to write transactions response", 500)
		}
	} else {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Unsupported method", 405)
	}
}

// SetupHandlers configures routes & handlers for the transaction server
func SetupHandlers(fs http.FileSystem) {
	http.HandleFunc("/transaction", handleTransaction)
	http.HandleFunc("/transaction/", handleTransaction)
	http.HandleFunc("/transactions", handleTransactions)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(fs)))
}
