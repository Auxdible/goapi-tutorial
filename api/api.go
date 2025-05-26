package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance parameters
type CoinBalanceParams struct {
	// Username is the user's username for whom the coin balance is requested
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// Code is the HTTP status code, usually 200 for success
	Code int

	// Balance is the user's coin balance
	Balance int64
}

// Error response
type Error struct {
	// Code is the HTTP status code, usually 400 or 500
	Code int

	// Message is a human-readable error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
