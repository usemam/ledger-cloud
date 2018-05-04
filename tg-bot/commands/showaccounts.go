package commands

// ShowAccountsCommand - command for getting current state of all registered accounts
type ShowAccountsCommand struct {
}

// Execute - Command interface implementation
func (cmd *ShowAccountsCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"show_accounts"}
}
