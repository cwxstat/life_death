package actu

import (
	"testing"
)

func Test_build_calc(t *testing.T) {

	b := newBuild()
	b.calc()
	if len(b.data) <= 0 {
		t.FailNow()
	}

}
