package bank

import (
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	trans := NewTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), "foo")
	json := trans.JSON()
	if json != "{\"id\":0,\"amount\":99,\"date\":\"2014-05-23T00:00:00Z\",\"note\":\"foo\"}" {
		t.Errorf("Bad JSON: %s", json)
	}
}

func TestSave(t *testing.T) {
	trans := NewTransaction(1, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), "foo")
	trans.Save()
	reloaded := LoadTransaction(trans.ID)

	if reloaded.ID != trans.ID {
		t.Errorf("Bad reloaded Id %d != %d", trans.ID, reloaded.ID)
	}
}

func TestSaveWithoutDate(t *testing.T) {
	trans := NewTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), "foo")
	trans.Date = nil
	err := trans.Save()

	if err == nil || err.Error() != "bank: cannot save Transaction without a Date" {
		t.Errorf("Expected an error on Save() without a Date")
	}
}

func TestSaveWithLongNote(t *testing.T) {
	note := ""
	for i := 0; i < 256; i++ {
		note += "a"
	}
	trans := NewTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC), note)
	err := trans.Save()

	if err == nil || err.Error() != "bank: transaction note is too long to save" {
		t.Errorf("Expected an error on Save() without a long note")
	}
}
