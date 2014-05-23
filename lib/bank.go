package bank

import (
    "fmt"
    "time"
    "encoding/json"
)

var (
    transactions = make([]*Transaction, 0)
)

type Transaction struct {
    Amount int
    Date *time.Time
}

func (t *Transaction) JSON() string {
    j, err := json.Marshal(t)
    if err != nil {
        fmt.Println("Failed to JSONify Transaction")
        return ""
    }
    return string(j)
}

func NewTransaction(amount int) *Transaction {
    transactionDate := time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC)
    return &Transaction{amount, &transactionDate}
}

func LoadTransaction(id int) *Transaction {
    trans := append(transactions, NewTransaction(105))
    return trans[0]
}
