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
