package bank

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var database *sql.DB

const maxNoteLength = 255

type transaction struct {
	ID     int        `json:"id"`
	Amount int        `json:"amount"`
	Date   *time.Time `json:"date"`
	Note   string     `json:"note"`
}

func newTransaction(amount int, transactionDate time.Time, note string) *transaction {
	return &transaction{
		Amount: amount,
		Date:   &transactionDate,
		Note:   note,
	}
}

// Instantiate and return a reference to a Transaction from the db
func loadTransaction(id int) *transaction {
	db := getDB()
	stmt, err := db.Prepare("SELECT amount, time, note FROM Transaction WHERE id=? AND deletionTime IS NULL")
	if err != nil {
		log.Printf("db.Prepare(): %v\n", err)
		return nil
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close statement: %v\n", err)
		}
	}()

	var amount int
	var transactionDate string
	var note []byte //nullable
	rowNotFound := stmt.QueryRow(id).Scan(&amount, &transactionDate, &note)
	if rowNotFound != nil {
		return nil
	}

	parsedTime, err := time.Parse("2006-01-02", transactionDate)
	if err != nil {
		log.Printf("failed to parse date %v: %v\n", transactionDate, err)
		return nil
	}
	trans := newTransaction(amount, parsedTime, string(note))
	trans.ID = id
	return trans
}

// Load multiple transactions
func loadTransactions() []*transaction {
	db := getDB()
	rows, err := db.Query("SELECT id, amount, time, note FROM Transaction WHERE deletionTime IS NULL ORDER BY time")
	if err != nil {
		log.Printf("failed to load transactions: %v\n", err)
		return nil
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close statement: %v\n", err)
		}
	}()

	var id, amount int
	var transactionDate string
	var note []byte //nullable
	result := []*transaction{}

	for rows.Next() {
		err := rows.Scan(&id, &amount, &transactionDate, &note)
		if err != nil {
			log.Printf("Scan(): %v\n", err)
			return nil
		}
		parsedTime, err := time.Parse("2006-01-02", transactionDate)
		if err != nil {
			log.Printf("time.Parse(): %v\n", err)
		}
		trans := newTransaction(amount, parsedTime, string(note))
		trans.ID = id
		result = append(result, trans)
	}
	return result
}

var (
	errNoTransactionDate = errors.New("cannot save Transaction without a Date")
	//FIXME: This is in bytes, not chars. This should be changed to be Unicode-compatible
	errNoteMaxLength        = fmt.Errorf("transaction note cannot exceed %d chars", maxNoteLength)
	errTransientTransaction = errors.New("cannot delete a transient Transaction")
)

// Persist a transaction
func (t *transaction) Save() error {
	if t.Date == nil {
		return errNoTransactionDate
	}
	if len(t.Note) > maxNoteLength {
		return errNoteMaxLength
	}

	db := getDB()
	stmt, err := db.Prepare("INSERT INTO Transaction (time, amount, note) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close statement: %v\n", err)
		}
	}()

	res, err := stmt.Exec(t.Date, t.Amount, t.Note)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// Set the lastId on the Transaction
	t.ID = int(lastID)
	return nil
}

// Delete a transaction
func (t *transaction) Delete() error {
	if t.ID == 0 {
		return errTransientTransaction
	}

	db := getDB()
	stmt, err := db.Prepare("UPDATE Transaction SET deletionTime=NOW() where id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.ID)
	return err
}

func getDB() *sql.DB {
	if database == nil {
		db, err := sql.Open("mysql", "bank:bank@/bank")
		if err != nil {
			log.Printf("sql.Open(): %v\n", err)
		}
		database = db
	}
	return database
}
