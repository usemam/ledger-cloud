package parser

import (
	"fmt"
	"strconv"
	str "strings"
	"time"

	p "github.com/araddon/dateparse"
	m "github.com/usemam/ledger-cloud/money"
)

// Token - scanned token
type Token struct {
	Value string
}

type moneyNotParsedError struct {
	Value string
}

func (e *moneyNotParsedError) Error() string {
	return fmt.Sprintf("%q cannot be parsed as Money value.", e.Value)
}

// Number - gets integer value of token
func (t *Token) Number() (int, error) {
	return strconv.Atoi(t.Value)
}

// Name - gets string value of token
func (t *Token) Name() string {
	return t.Value
}

// Date - gets date value of token
func (t *Token) Date() (time.Time, error) {
	return p.ParseAny(t.Value)
}

// Money - gets money value of token
func (t *Token) Money() (m.Money, error) {
	split := str.Split(t.Value, ".")
	whole, err := strconv.Atoi(split[0])
	if err != nil {
		return m.Zero(), err
	}
	if len(split) == 1 {
		return m.New(uint(whole), 0), nil
	}
	if len(split) == 2 {
		part, err := strconv.Atoi(split[1])
		if err != nil {
			return m.Zero(), err
		}
		return m.New(uint(whole), uint(part)), nil
	}
	return m.Zero(), &moneyNotParsedError{t.Value}
}
