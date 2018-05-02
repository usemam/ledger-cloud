package configuration

import (
	"os"
)

// Configuration - structure to hold current env. settings
type Configuration struct {
	Port              string
	TransactionAPIURL string
	RedisURL          string
}

// New - factory method to read current configuration
func New() *Configuration {
	return &Configuration{
		Port:              os.Getenv("PORT"),
		TransactionAPIURL: os.Getenv("TRN_API_URL"),
		RedisURL:          os.Getenv("REDIS_URL"),
	}
}
