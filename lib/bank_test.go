package bank

import (
	"testing"
	"time"
)

func TestJSON(t *testing.T) {
	trans := NewTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC))
	json := trans.JSON()
	if json != "{\"Id\":0,\"Amount\":99,\"Date\":\"2014-05-23T00:00:00Z\"}" {
		t.Errorf("Bad JSON: %s", json)
	}
}

func TestSave(t *testing.T) {
	trans := NewTransaction(1, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC))
	trans.Save()
	reloaded := LoadTransaction(trans.Id)

	if reloaded.Id != trans.Id {
		t.Errorf("Bad reloaded Id %d != %d", trans.Id, reloaded.Id)
	}
}
