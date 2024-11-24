package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"wallet-service/internal/api"
	"wallet-service/internal/db"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/wallet", api.HandlePostWallet).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{walletId}", api.HandleGetWallet).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
