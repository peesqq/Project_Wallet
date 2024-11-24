package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      uuid.UUID `json:"walletId"`
	Balance int64     `json:"balance"`
	Created time.Time `json:"created"`
}

type OperationType string

const (
	Deposit  OperationType = "DEPOSIT"
	Withdraw OperationType = "WITHDRAW"
)

type Operation struct {
	WalletID      uuid.UUID     `json:"walletId"`
	OperationType OperationType `json:"operationType"`
	Amount        int64         `json:"amount"`
	Timestamp     time.Time     `json:"timestamp"`
}
