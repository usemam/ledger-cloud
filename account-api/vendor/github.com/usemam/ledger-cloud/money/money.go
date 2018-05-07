package money

import (
	"fmt"
	str "strings"
)

// Money - money abstraction
type Money struct {
	Whole uint
	Part  uint
}

// New - factory method
func New(whole uint, part uint) Money {
	return Money{whole, part}
}

// Zero - factory method, inits with zero values
func Zero() Money {
	return Money{}
}

// Gt - determines if x > y
func (x Money) Gt(y Money) bool {
	return x.Whole > y.Whole || x.Whole == y.Whole && x.Part > y.Part
}

// Add - summing two Money values
func (x Money) Add(y Money) Money {
	partSum := x.Part + y.Part
	return Money{x.Whole + y.Whole + partSum/100, partSum % 100}
}

// Sub - subtracting two Money values
func (x Money) Sub(y Money) Money {
	var (
		whole uint
		part  uint
	)
	if x.Part < y.Part {
		whole = x.Whole - y.Whole - 1
		part = 100 + x.Part - y.Part
	} else {
		whole = x.Whole - y.Whole
		part = x.Part - y.Part
	}
	return Money{whole, part}
}

func (x Money) String() string {
	if x.Whole == 0 {
		return fmt.Sprintf("$0.%02d", x.Part)
	}

	builder := &str.Builder{}
	builder.WriteString("$")
	parts := getThousandParts(x.Whole)
	for i := 0; i < len(parts); i++ {
		if i > 0 {
			builder.WriteString(fmt.Sprintf(",%03d", parts[i]))
		} else {
			builder.WriteString(fmt.Sprintf("%d", parts[i]))
		}
	}
	builder.WriteString(fmt.Sprintf(".%02d", x.Part))
	return builder.String()
}

func getThousandParts(n uint) []uint {
	a := make([]uint, 0)
	for i := n; i > 0; i /= 1000 {
		a = append(a, i%1000)
	}
	return reverseArr(a)
}

func reverseArr(a []uint) []uint {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}
