package commands

import (
	"fmt"
)

// Command - common interface for all bot's command handlers
type Command interface {
	Execute(args string) (string, error)
}

// CreateCommand - factory method to create command instance based on provided name
func CreateCommand(name string) (Command, error) {
	switch name {
	case "new_account":
		return nil, nil
	case "transfer":
		return nil, nil
	case "debit":
		return nil, nil
	case "credit":
		return nil, nil
	case "show_accounts":
		return &ShowAccountsCommand{}, nil
	case "last":
		return nil, nil
	case "total":
		return nil, nil
	default:
		return nil, &unrecognizedCommandError{name}
	}
}

// NotImplementedCommandError - command is not implemented yet
type NotImplementedCommandError struct {
	CommandName string
}

func (err *NotImplementedCommandError) Error() string {
	return fmt.Sprintf("Command %q is not yet implemented.", err.CommandName)
}

type unrecognizedCommandError struct {
	CommandName string
}

func (err *unrecognizedCommandError) Error() string {
	return fmt.Sprintf("Unrecognized command - %q.", err.CommandName)
}
