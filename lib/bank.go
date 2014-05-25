package bank

import (
    "fmt"
    "time"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
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
    db, err := sql.Open("mysql", "root@/bank")
    if err != nil {
        fmt.Println("Error connecting to MySQL")
        fmt.Println(err)
        return nil
    }
    defer db.Close()

    stmt, err := db.Prepare("SELECT amount, time FROM Transaction WHERE id=?")
    if err != nil {
        fmt.Println("Error preparing statement")
        fmt.Println(err)
        return nil
    }
    defer stmt.Close()

    var amount int
    var time string
    stmt.QueryRow(id).Scan(&amount, &time)
    fmt.Println(fmt.Sprintf("amount:%d time:%s", amount, time))

    trans := append(transactions, NewTransaction(105))
    return trans[0]
}
