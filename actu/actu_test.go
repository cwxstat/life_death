package actu

import (
	"testing"
)

func TestActu_CSV(t *testing.T) {

	a := NewActu()
	if len(a.Data()) <= 10 {
		t.FailNow()
	}
}
