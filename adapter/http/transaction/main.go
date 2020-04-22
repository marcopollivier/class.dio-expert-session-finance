package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/marcopollivier/dio-expert-session-pre-class/model/transaction"
	"io/ioutil"
	"net/http"
	"time"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	layout := "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45:26")
	paidElectricityBill, _ := time.Parse(layout, "2020-04-12T22:00:00")
	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:       "Salário",
			Amount:      1200.0,
			Type:        0,
			Installment: 1,
			CreatedAt:   salaryReceived,
		},
		transaction.Transaction{
			Title:     "Conta de luz",
			Amount:    100.0,
			Type:      1,
			Installment: 1,
			CreatedAt: paidElectricityBill,
		},
		transaction.Transaction{
			Title:     "Compra telefone celular",
			Amount:    999.99,
			Type:      1,
			Installment: 10,
			CreatedAt: paidElectricityBill,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &res)

	fmt.Println(res)
}