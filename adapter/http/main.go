package http

import (
	"net/http"

	"github.com/LaisaCCAndrade/expert-session-finance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.ListenAndServe(":8080", nil)
}
