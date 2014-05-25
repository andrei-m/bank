package bank

import (
    "testing"
    "time"
)

func TestJSON(t *testing.T) {
    trans := NewTransaction(99, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC))
    json := trans.JSON()
    if json != "{\"Amount\":99,\"Date\":\"2014-05-23T00:00:00Z\"}" {
        t.Errorf("Bad JSON: %s", json)
    }
}

func TestSave(t *testing.T) {
    trans := NewTransaction(1, time.Date(2014, time.May, 23, 0, 0, 0, 0, time.UTC))
    trans.Save()
}
    
