package money

import (
	"fmt"
	tst "testing"
)

func TestStringFormatting(t *tst.T) {
	m := map[string]Money{
		"$0.00":         Money{0, 0},
		"$100.00":       Money{100, 0},
		"$1,000.00":     Money{1000, 0},
		"$101,010.10":   Money{101010, 10},
		"$1,000,000.05": Money{1000000, 5},
	}

	for k, v := range m {
		s := v.String()
		if s != k {
			t.Log(fmt.Sprintf("Expected %s but was %s", k, s))
			t.Fail()
		}
	}
}
