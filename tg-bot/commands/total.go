package commands

// TotalCommand - command to show total balance for some or all categories
type TotalCommand struct {
	UserID string
}

// Execute - Command interface implementation
// Usage: /total [date_from] [date_to] [category_name]
func (cmd *TotalCommand) Execute(args string) (string, error) {
	return "", &NotImplementedCommandError{"total"}
}
