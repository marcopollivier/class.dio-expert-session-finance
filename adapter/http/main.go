package http

import (
	"github.com/marcopollivier/dio-expert-session-pre-class/adapter/http/actuator"
	"github.com/marcopollivier/dio-expert-session-pre-class/adapter/http/transaction"
	"net/http"
)

func Init() error{
	http.HandleFunc("/health", actuator.Health)

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	return http.ListenAndServe(":8080", nil)
}