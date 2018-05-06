package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/usemam/ledger-cloud/transaction-api/configuration"
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
