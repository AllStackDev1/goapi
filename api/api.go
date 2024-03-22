package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParam struct {
	Username string
}

// Coin Balance Response
type CoinBalanceBalance struct {
	// Success Cocd, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error Code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHander = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)