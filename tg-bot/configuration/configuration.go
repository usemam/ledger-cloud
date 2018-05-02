package configuration

import (
	"os"
)

// Configuration - structure to hold current env. settings
type Configuration struct {
	Port              string
	BotToken          string
	BotURL            string
	TransactionAPIURL string
	AccountAPIURL     string
}

// New - factory method to read current configuration
func New() *Configuration {
	return &Configuration{
		Port:              os.Getenv("PORT"),
		BotToken:          os.Getenv("TOKEN"),
		BotURL:            os.Getenv("URL"),
		TransactionAPIURL: os.Getenv("TRN_API_URL"),
		AccountAPIURL:     os.Getenv("ACNT_API_URL"),
	}
}
