package commands

// CreditCommand - command to deposit money to account from earning category
type CreditCommand struct {
}

// Execute - Command interface implementation
func (cmd *CreditCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"credit"}
}
