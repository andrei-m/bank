package main

import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/andrei-m/bank/lib"
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
        fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
    } else {
        w.Header().Set("Allow", "GET, POST")
        http.Error(w, "Unsupported method", 405)
    }
}

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/transaction/", handleTransaction)
    http.ListenAndServe(":1337", nil)
}

