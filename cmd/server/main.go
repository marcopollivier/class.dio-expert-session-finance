package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)

	_ = http.ListenAndServe(":8080", nil)
}


type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"` //0. entrada 1. saida
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	layout := "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45:26")
	paidElectricityBill, _ := time.Parse(layout, "2020-04-12T22:00:00")
	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
		Transaction{
			Title:     "Conta de luz",
			Amount:    100.0,
			Type:      1,
			CreatedAt: paidElectricityBill,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}