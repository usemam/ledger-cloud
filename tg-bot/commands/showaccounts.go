package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	str "strings"

	dc "github.com/usemam/ledger-cloud/data-contract"
	configuration "github.com/usemam/ledger-cloud/tg-bot/configuration"
)

// ShowAccountsCommand - command for getting current state of all registered accounts
type ShowAccountsCommand struct {
	UserID string
}

// Execute - Command interface implementation
func (cmd *ShowAccountsCommand) Execute(args string) (string, error) {
	cfg := configuration.New()
	url := fmt.Sprintf(cfg.AccountAPIURL+"state/%s", cmd.UserID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var accounts []dc.Account
	if err := json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return "", err
	}

	replyBuilder := &str.Builder{}
	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		if replyBuilder.Len() > 0 {
			replyBuilder.WriteString("\n")
		}
		replyBuilder.WriteString(fmt.Sprintf("%d. %s - $%f", i, account.Name, account.Balance))
	}

	return replyBuilder.String(), nil
}
