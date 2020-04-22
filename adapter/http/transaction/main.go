package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/marcopollivier/dio-expert-session-pre-class/model/transaction"
	"github.com/marcopollivier/dio-expert-session-pre-class/util"
	"io/ioutil"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:       "Salário",
			Amount:      1200.0,
			Type:        0,
			Installment: 1,
			CreatedAt:   util.StringToTime("2020-04-05T11:45:26"),
		},
		transaction.Transaction{
			Title:     "Conta de luz",
			Amount:    100.0,
			Type:      1,
			Installment: 1,
			CreatedAt: util.StringToTime("2020-04-12T22:00:00"),
		},
		transaction.Transaction{
			Title:     "Compra telefone celular",
			Amount:    999.99,
			Type:      1,
			Installment: 10,
			CreatedAt: util.StringToTime("2020-04-20T11:00:26"),
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