package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

// Модель операции
type Operation struct {
	WalletID      string `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount"`
}

func TestMain(m *testing.M) {
	// Загружаем переменные окружения из .env.test
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}

	// Запускаем тесты
	m.Run()
}

// Тестирование депозита
func TestHandlePostWalletDeposit(t *testing.T) {
	op := Operation{
		WalletID:      "cb7a4b44-8dc2-4fec-9277-61b1fcb5a264",
		OperationType: "DEPOSIT",
		Amount:        500,
	}
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatal("Error marshalling request body:", err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/wallet", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error executing request:", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Тестирование снятия средств
func TestHandlePostWalletWithdraw(t *testing.T) {
	op := Operation{
		WalletID:      "cb7a4b44-8dc2-4fec-9277-61b1fcb5a264",
		OperationType: "WITHDRAW",
		Amount:        300,
	}
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatal("Error marshalling request body:", err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/wallet", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error executing request:", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Тестирование получения баланса кошелька
func TestHandleGetWallet(t *testing.T) {
	url := "http://localhost:8080/api/v1/wallets/cb7a4b44-8dc2-4fec-9277-61b1fcb5a264"

	// Выполнение GET запроса
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal("Error executing GET request:", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
