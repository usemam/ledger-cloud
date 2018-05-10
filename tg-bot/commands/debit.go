package commands

// DebitCommand - command to withdraw money from account to spending category
type DebitCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /debit account_name amount category_name [date]
func (cmd *DebitCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"debit"}
}
