package commands

// NewAccountCommand - command to create new account
type NewAccountCommand struct {
}

// Execute - Command interface implementation
func (cmd *NewAccountCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"new_account"}
}
