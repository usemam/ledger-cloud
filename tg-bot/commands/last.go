package commands

import (
	p "github.com/usemam/ledger-cloud/tg-bot/commands/arg-parser"
)

// LastCommand - command to show last <n> operations
type LastCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /last n [account_name]
func (cmd *LastCommand) Execute(args string) (string, error) {
	scan := p.NewScanner(args).Number("n").Name("account_name").Optional()
	return "", &NotImplementedCommandError{"last"}
}
