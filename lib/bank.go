package bank

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var database *sql.DB

const maxNoteLength = 255

type Transaction struct {
	Id     int
	Amount int
	Date   *time.Time
	Note   string
}

func (t *Transaction) JSON() string {
	j, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to JSONify Transaction")
		return ""
	}
	return string(j)
}

func NewTransaction(amount int, transactionDate time.Time, note string) *Transaction {
	return &Transaction{0, amount, &transactionDate, note}
}

// Instantiate and return a reference to a Transaction from the db
func LoadTransaction(id int) *Transaction {
	db := getDB()
	stmt, err := db.Prepare("SELECT amount, time, note FROM Transaction WHERE id=?")
	if err != nil {
		fmt.Println("Error preparing statement")
		fmt.Println(err)
		return nil
	}
	defer stmt.Close()

	var amount int
	var transactionDate string
	var note []byte //nullable
	rowNotFound := stmt.QueryRow(id).Scan(&amount, &transactionDate, &note)
	if rowNotFound != nil {
		return nil
	}

	parsedTime, err := time.Parse("2006-01-02", transactionDate)
	trans := NewTransaction(amount, parsedTime, string(note))
	trans.Id = id
	return trans
}

// Load multiple transactions
func LoadTransactions() []*Transaction {
	db := getDB()
	rows, err := db.Query("SELECT id, amount, time, note FROM Transaction ORDER BY time")
	if err != nil {
		fmt.Println("Error retrieving transactions")
		fmt.Println(err)
	}
	defer rows.Close()

	var id, amount int
	var transactionDate string
	var note []byte //nullable
	result := make([]*Transaction, 0)

	for rows.Next() {
		err := rows.Scan(&id, &amount, &transactionDate, &note)
		if err != nil {
			fmt.Println("Failed to scan")
			fmt.Println(err)
		} else {
			parsedTime, _ := time.Parse("2006-01-02", transactionDate)

			trans := NewTransaction(amount, parsedTime, string(note))
			trans.Id = id
			result = append(result, trans)
		}
	}

	return result
}

// Persist a transaction
func (t *Transaction) Save() error {
	if t.Date == nil {
		return errors.New("bank: cannot save Transaction without a Date")
	}
	if len(t.Note) > maxNoteLength {
		return errors.New("bank: transaction note is too long to save")
	}

	db := getDB()
	stmt, err := db.Prepare("INSERT INTO Transaction (time, amount, note) VALUES (?, ?, ?)")
	if err != nil {
		return errors.New("bank: error preparing statement for Save()")
	}
	defer stmt.Close()

	res, err := stmt.Exec(t.Date, t.Amount, t.Note)
	if err != nil {
		return errors.New("bank: error executing statement for Save()")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return errors.New("bank: retrieving last id for Save()")
	}

	// Set the lastId on the Transaction
	t.Id = int(lastId)
	return nil
}

func getDB() *sql.DB {
	if database == nil {
		db, err := sql.Open("mysql", "bank:bank@/bank")
		if err != nil {
			fmt.Println("Error connecting to MySQL")
			fmt.Println(err)
		}
		database = db
	}

	return database
}
