package commands

// DebitCommand - command to withdraw money from account to spending category
type DebitCommand struct {
}

// Execute - Command interface implementation
func (cmd *DebitCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"debit"}
}
