package main

import (
	"net/http"

	configuration "github.com/usemam/ledger-cloud/tg-bot/configuration"
)

func main() {
	cfg := configuration.New()
	go http.ListenAndServe(":"+cfg.Port, nil)
}
