package transaction

import "time"

//type
//	0. entrada
//	1. saida
type Transaction struct {
	Title       string    `json:"title"`
	Amount      float32   `json:"amount"`
	Type        int       `json:"type"`
	Installment int       `json:"installment"`
	CreatedAt   time.Time `json:"created_at"`
}

type Transactions []Transaction
