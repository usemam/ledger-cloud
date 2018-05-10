package commands

// NewAccountCommand - command to create new account
type NewAccountCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /new_account account_name balance [max_credit]
func (cmd *NewAccountCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"new_account"}
}
