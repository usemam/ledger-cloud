package commands

// TransferCommand - command to transfer money between accounts
type TransferCommand struct {
	UserID string
}

// Execute - Command interface implementation
func (cmd *TransferCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"transfer"}
}
