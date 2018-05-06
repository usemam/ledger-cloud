package commands

import (
	"fmt"
)

// Command - common interface for all bot's command handlers
type Command interface {
	Execute(args string) (string, error)
}

// CreateCommand - factory method to create command instance based on provided name
func CreateCommand(cmdName string, userID string) (Command, error) {
	switch cmdName {
	case "new_account":
		return &NewAccountCommand{userID}, nil
	case "transfer":
		return &TransferCommand{userID}, nil
	case "debit":
		return &DebitCommand{userID}, nil
	case "credit":
		return &CreditCommand{userID}, nil
	case "show_accounts":
		return &ShowAccountsCommand{userID}, nil
	case "last":
		return &LastCommand{userID}, nil
	case "total":
		return &TotalCommand{userID}, nil
	default:
		return nil, &unrecognizedCommandError{cmdName}
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
