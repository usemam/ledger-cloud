package parser

import (
	"strconv"
)

// Token - scanned token
type Token struct {
	Value string
}

// Number - gets integer value of token
func (t *Token) Number() (int, error) {
	return strconv.Atoi(t.Value)
}

// Name - gets string value of token
func (t *Token) Name() string {
	return t.Value
}
