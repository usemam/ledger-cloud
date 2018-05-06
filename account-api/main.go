package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/usemam/ledger-cloud/account-api/configuration"
	"github.com/usemam/ledger-cloud/account-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/state/{user_id}", handlers.ShowAccountsHandler)

	cfg := configuration.New()
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
