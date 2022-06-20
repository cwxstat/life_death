package loader

import (
	"testing"
)

func Test_build_Calc(t *testing.T) {
	b := NewLoad()
	b.Calc()
	if len(b.numbers) < 100 {
		t.FailNow()
	}

}
