package configuration

import (
	"os"
)

// Configuration - structure to hold current env. settings
type Configuration struct {
	Port          string
	AccountAPIURL string
	MongoURL      string
}

// New - factory method to read current configuration
func New() *Configuration {
	return &Configuration{
		Port:          os.Getenv("PORT"),
		AccountAPIURL: os.Getenv("ACNT_API_URL"),
		MongoURL:      os.Getenv("MONGO_URL"),
	}
}
