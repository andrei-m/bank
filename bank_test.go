package bank

import (
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	trans := newTransaction(1, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), "foo")
	if err := trans.Save(); err != nil {
		t.Errorf("failed to save transaction: %v", err)
	}
	reloaded := loadTransaction(trans.ID)

	if reloaded.ID != trans.ID {
		t.Errorf("Bad reloaded Id %d != %d", trans.ID, reloaded.ID)
	}
}

func TestSaveWithoutDate(t *testing.T) {
	trans := newTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), "foo")
	trans.Date = nil
	err := trans.Save()

	if err != errNoTransactionDate {
		t.Errorf("Expected an error on Save() without a Date")
	}
}

func TestSaveWithLongNote(t *testing.T) {
	note := ""
	for i := 0; i < 256; i++ {
		note += "a"
	}
	trans := newTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), note)
	err := trans.Save()

	if err != errNoteMaxLength {
		t.Errorf("Expected an error on Save() without a long note")
	}
}
