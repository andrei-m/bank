package bank

import "testing"

func TestJSON(t *testing.T) {
    trans := NewTransaction(99)
    json := trans.JSON()
    if json != "{\"Amount\":99}" {
        t.Errorf("Bad JSON: %s", json)
    }
}
    
