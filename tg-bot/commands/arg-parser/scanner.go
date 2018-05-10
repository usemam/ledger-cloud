package parser

import (
	str "strings"
)

const (
	// Number - positive number parameter
	Number = iota
	// Name - string parameter
	Name = iota
	// Date - date parameter
	Date = iota
	// Money - money parameter
	Money = iota
)

// ParamDefinition - parameter definition
type ParamDefinition struct {
	Type       int
	IsOptional bool
	Name       string
}

// Scanner - arguments parser
type Scanner struct {
	InputPtr         int
	ParamPtr         int
	Input            []string
	ParamDefinitions map[int]*ParamDefinition
}

// NewScanner - factory method for Scanner
func NewScanner(input string) *Scanner {
	return &Scanner{
		Input:            str.Split(input, " "),
		ParamDefinitions: make(map[int]*ParamDefinition),
		ParamPtr:         1,
	}
}

// Optional - marks last added parameter as optional
func (s *Scanner) Optional() *Scanner {
	param, ok := s.ParamDefinitions[len(s.ParamDefinitions)]
	if ok {
		param.IsOptional = true
	}
	return s
}

// Number - adds number parameter
func (s *Scanner) Number(name string) *Scanner {
	count := len(s.ParamDefinitions)
	s.ParamDefinitions[count+1] = &ParamDefinition{Type: Number, Name: name}
	return s
}

// Name - adds name parameter
func (s *Scanner) Name(paramName string) *Scanner {
	count := len(s.ParamDefinitions)
	s.ParamDefinitions[count+1] = &ParamDefinition{Type: Name, Name: paramName}
	return s
}

// Date - adds date parameter
func (s *Scanner) Date(name string) *Scanner {
	count := len(s.ParamDefinitions)
	s.ParamDefinitions[count+1] = &ParamDefinition{Type: Date, Name: name}
	return s
}

// Money - adds money parameter
func (s *Scanner) Money(name string) *Scanner {
	count := len(s.ParamDefinitions)
	s.ParamDefinitions[count+1] = &ParamDefinition{Type: Money, Name: name}
	return s
}
