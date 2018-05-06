package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	dc "github.com/usemam/ledger-cloud/account-api/datacontract"
)

// ShowAccountsHandler - handles /state/{user_id} query
func ShowAccountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	userID := vars["user_id"]
	accounts := [1]dc.Account{
		dc.Account{
			UserID:      userID,
			Name:        "Tinkoff Black",
			Balance:     1000.0,
			MaxCredit:   0.0,
			LastUpdated: time.Now().Unix(),
		},
	}
	json.NewEncoder(w).Encode(accounts)
}
