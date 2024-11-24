package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"wallet-service/internal/db"
	"wallet-service/internal/models"
)

var mu sync.Mutex

func HandlePostWallet(w http.ResponseWriter, r *http.Request) {
	var op models.Operation

	if err := json.NewDecoder(r.Body).Decode(&op); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if op.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	var newBalance int64
	row := db.DB.QueryRow(
		r.Context(),
		"SELECT balance FROM wallets WHERE id=$1",
		op.WalletID,
	)

	var currentBalance int64
	if err := row.Scan(&currentBalance); err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	switch op.OperationType {
	case models.Deposit:
		newBalance = currentBalance + op.Amount
	case models.Withdraw:
		if currentBalance < op.Amount {
			http.Error(w, "Insufficient balance", http.StatusBadRequest)
			return
		}
		newBalance = currentBalance - op.Amount
	default:
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(
		r.Context(),
		"UPDATE wallets SET balance=$1 WHERE id=$2",
		newBalance, op.WalletID,
	)

	if err != nil {
		http.Error(w, "Failed to update wallet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["walletId"]

	id, err := uuid.Parse(walletID)
	if err != nil {
		http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
		return
	}

	var balance int64
	row := db.DB.QueryRow(r.Context(), "SELECT balance FROM wallets WHERE id=$1", id)

	if err := row.Scan(&balance); err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"walletId": id,
		"balance":  balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
