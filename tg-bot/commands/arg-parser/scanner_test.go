package parser

import (
	str "strings"
	tst "testing"
)

func TestNewScanner_InputSplitBySpaces(t *tst.T) {
	input := "test input"
	scan := NewScanner(input)
	split := str.Split(input, " ")
	if len(scan.Input) != 2 || scan.Input[0] != split[0] || scan.Input[1] != split[1] {
		t.Error("Scanner's input not initialized properly")
	}
}

func TestNumber(t *tst.T) {
	name := "n"
	scan := NewScanner("").Number(name)
	param, ok := scan.ParamDefinitions[1]
	if !ok || param.Name != name || param.Type != Number {
		t.Error("Adding number parameter has failed.")
	}
}

func TestName(t *tst.T) {
	name := "account"
	scan := NewScanner("").Name(name)
	param, ok := scan.ParamDefinitions[1]
	if !ok || param.Name != name || param.Type != Name {
		t.Error("Adding name parameter has failed.")
	}
}

func TestDate(t *tst.T) {
	name := "date"
	scan := NewScanner("").Date(name)
	param, ok := scan.ParamDefinitions[1]
	if !ok || param.Name != name || param.Type != Date {
		t.Error("Adding date parameter has failed.")
	}
}

func TestMoney(t *tst.T) {
	name := "amount"
	scan := NewScanner("").Money(name)
	param, ok := scan.ParamDefinitions[1]
	if !ok || param.Name != name || param.Type != Money {
		t.Error("Adding money parameter has failed.")
	}
}

func TestOptional(t *tst.T) {
	name := "optional"
	scan := NewScanner("").Name(name).Optional()
	param, ok := scan.ParamDefinitions[1]
	if !ok || !param.IsOptional || param.Name != name {
		t.Error("Adding optional parameter has failed.")
	}
}

func TestScanner_AddManyParams(t *tst.T) {
	name1, name2, name3 := "p1", "p2", "p3"
	scan := NewScanner("").Date(name1).Money(name2).Name(name3)
	param, ok := scan.ParamDefinitions[1]
	if !ok || param.Name != name1 {
		t.Error("Adding first parameter has failed.")
	}
	param, ok = scan.ParamDefinitions[2]
	if !ok || param.Name != name2 {
		t.Error("Adding second parameter has failed.")
	}
	param, ok = scan.ParamDefinitions[3]
	if !ok || param.Name != name3 {
		t.Error("Adding third parameter has failed.")
	}
}
