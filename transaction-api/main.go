package main

import (
	"net/http"

	configuration "github.com/usemam/ledger-cloud/transaction-api/configuration"
)

func main() {
	cfg := configuration.New()
	go http.ListenAndServe(":"+cfg.Port, nil)
}
