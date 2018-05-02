package main

import (
	"net/http"

	configuration "github.com/usemam/ledger-cloud/account-api/configuration"
)

func main() {
	cfg := configuration.New()
	go http.ListenAndServe(":"+cfg.Port, nil)
}
