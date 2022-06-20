package actu

import (
	"fmt"
	"testing"
)

func TestActu_CSV(t *testing.T) {

	a := NewActu()
	if len(a.Data()) <= 10 {
		t.FailNow()
	}
}

func TestActu_Bucket(t *testing.T) {
	a := NewActu()
	a.Bucket(15)
}

func TestActu_Range(t *testing.T) {
	a := NewActu()
	result := a.Range(0, 65)
	fmt.Println(result)
}
