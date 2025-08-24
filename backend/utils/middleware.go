package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	sqlc "github.com/triyaambak/CRM/internal/sqlc_db"
)

func WriteResponseMiddleware(w http.ResponseWriter, leads []sqlc.Lead, errors []error, success int, failure int) {
	w.Header().Set("Content-Type", "application/json")

	resp := struct {
		Leads   []sqlc.Lead `json:"leads"`
		Errors  []error     `json:"errors"`
		Success int         `json:"success"`
		Failure int         `json:"failure"`
	}{
		Leads:   leads,
		Errors:  errors,
		Success: success,
		Failure: failure,
	}

	w.WriteHeader(http.StatusOK)

	if success == 0 && failure > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode JSON response in GetAllLeads function %v", err)
		return
	}

}
