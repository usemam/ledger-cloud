package commands

// TransferCommand - command to transfer money between accounts
type TransferCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /transfer account_from amount account_to [date]
func (cmd *TransferCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"transfer"}
}
