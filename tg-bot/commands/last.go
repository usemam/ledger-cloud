package commands

// LastCommand - command to show last <n> operations
type LastCommand struct {
	UserID string
}

// Execute - Command interface implementation
func (cmd *LastCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"last"}
}
