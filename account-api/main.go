package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	configuration "github.com/usemam/ledger-cloud/account-api/configuration"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", testHandler)

	cfg := configuration.New()
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
