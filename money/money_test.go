package money

import (
	tst "testing"
)

func TestString_Formatting(t *tst.T) {
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
			t.Errorf("Expected %s but was %s", k, s)
		}
	}
}

func TestAdd_Zero(t *tst.T) {
	var (
		whole uint = 5
		part  uint = 20
	)
	m := Money{whole, part}.Add(Money{})
	if m.Whole != whole || m.Part != part {
		t.Error("Adding zero resulted amounts different from original.")
	}
}

func TestSub_Zero(t *tst.T) {
	var (
		whole uint = 5
		part  uint = 20
	)
	m := Money{whole, part}.Sub(Money{})
	if m.Whole != whole || m.Part != part {
		t.Error("Subtracting zero resulted amounts different from original.")
	}
}

func TestAdd_PartCarryToWhole(t *tst.T) {
	var (
		whole1 uint = 5
		part1  uint = 20
		whole2 uint = 1
		part2  uint = 80
	)
	m := Money{whole1, part1}.Add(Money{whole2, part2})
	if m.Whole != (whole1+whole2+1) || m.Part != 0 {
		t.Error("Part carry when adding two values is incorrect.")
	}
}

func TestSub_PartCarryToWhole(t *tst.T) {
	var (
		whole1 uint = 5
		part1  uint = 20
		whole2 uint = 1
		part2  uint = 80
	)
	m := Money{whole1, part1}.Sub(Money{whole2, part2})
	if m.Whole != (whole1-whole2-1) || m.Part != (100+part1-part2) {
		t.Error("Part carry when adding two values is incorrect.")
	}
}

func TestGt_Zero(t *tst.T) {
	res := Money{}.Gt(Money{})
	if res {
		t.Error("Zero is not greater than zero.")
	}
}

func TestGt_True(t *tst.T) {
	res1, res2 := Money{5, 0}.Gt(Money{4, 0}), Money{5, 20}.Gt(Money{5, 0})
	if !res1 || !res2 {
		t.Fail()
	}
}

func TestGt_False(t *tst.T) {
	res1, res2 := Money{4, 0}.Gt(Money{5, 0}), Money{5, 0}.Gt(Money{5, 20})
	if res1 || res2 {
		t.Fail()
	}
}
