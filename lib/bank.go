package bank

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Transaction struct {
	Id     int
	Amount int
	Date   *time.Time
}

func (t *Transaction) JSON() string {
	j, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to JSONify Transaction")
		return ""
	}
	return string(j)
}

func NewTransaction(amount int, transactionDate time.Time) *Transaction {
	return &Transaction{0, amount, &transactionDate}
}

// Instantiate and return a reference to a Transaction from the db
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
	var transactionDate string
	rowNotFound := stmt.QueryRow(id).Scan(&amount, &transactionDate)
	if rowNotFound != nil {
		return nil
	}

	parsedTime, err := time.Parse("2006-01-02", transactionDate)
	trans := NewTransaction(amount, parsedTime)
	trans.Id = id
	return trans
}

// Load multiple transactions
//TODO: left off here
func LoadTransactions() []*Transaction {
    db, err := sql.Open("mysql", "root@/bank")
	if err != nil {
		fmt.Println("Error connecting to MySQL")
		fmt.Println(err)
		return nil
	}
	defer db.Close()

    rows, err := db.Query("SELECT id, amount, time FROM Transaction")
    if err != nil {
        fmt.Println("Error retrieving transactions")
        fmt.Println(err)
    }
    defer rows.Close()

    var id, amount int
    var transactionDate string

    for rows.Next() {
        err := rows.Scan(&id, &amount, &transactionDate)
        if err != nil {
            fmt.Println("Failed to scan")
            fmt.Println(err)
            parsedTime, _ := time.Parse("2006-01-02", transactionDate)
            
            trans := NewTransaction(amount, parsedTime)
            trans.Id = id
        }
    }

    result := make([]*Transaction, 0)
    return result
}

// Persist a transaction
func (t *Transaction) Save() {
	db, err := sql.Open("mysql", "root@/bank")
	if err != nil {
		fmt.Println("Error connecting to MySQL")
		fmt.Println(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Transaction (time, amount) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement")
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.Date, t.Amount)
	if err != nil {
		fmt.Println("Error executing statement")
		fmt.Println(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error retrieving last id")
		fmt.Println(err)
	}

	// Set the lastId on the Transaction
	t.Id = int(lastId)
}
