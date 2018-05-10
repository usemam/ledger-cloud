package commands

// CreditCommand - command to deposit money to account from earning category
type CreditCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /credit category_name amount account_name [date]
func (cmd *CreditCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"credit"}
}
