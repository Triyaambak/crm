package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sqlc "github.com/triyaambak/CRM/internal/sqlc_db"
	utils "github.com/triyaambak/CRM/utils"

	"github.com/gorilla/mux"
	"github.com/triyaambak/CRM/config"
)

type Controller struct{}

type LeadStruct struct {
	Name      string         `json:"name"`
	Phone     sql.NullString `json:"phone"`
	CreatedAt time.Time      `json:"createdat"`
}

func (c *Controller) ServeHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Or one more way to write on the Response object
		// w.Write[]{"Welcome to home page"}
		fmt.Fprintln(w, "Welcome to home page")
	}
}

func (c *Controller) AddLead(dbcfg *config.DbConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data LeadStruct

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong while decoding json to struct in AddLead method %v", err)
			return
		}

		h := utils.Helper{}

		lead, err := h.CheckInputData(data.Name, data.Phone, data.CreatedAt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
		}

		if err := dbcfg.Query.CreateLead(r.Context(), lead); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not add data to the database %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Data added to the database successfully")

	}
}

func (c *Controller) GetLead(dbcfg *config.DbConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := mux.Vars(r)["id"]

		if !ok {
			http.Error(w, "No ID param found ", http.StatusBadRequest)
			return
		}

		leadChan := make(chan sqlc.Lead)
		errChan := make(chan error)

		dbcfg.GetDataAsync(leadChan, errChan, r, fmt.Sprintf("SELECT * FROM LEADS WHERE ID=%s", userID))

		var leads []sqlc.Lead
		var errors []error
		var success int = 0
		var failure int = 0

		for leadChan != nil || errChan != nil {
			select {
			case lead, ok := <-leadChan:
				if !ok {
					leadChan = nil
					continue
				}

				leads = append(leads, lead)
				success++

			case err, ok := <-errChan:
				if !ok {
					errChan = nil
					continue
				}

				errors = append(errors, err)
				failure++
			}
		}

		utils.WriteResponseMiddleware(w, leads, errors, success, failure)

	}
}

func (c *Controller) GetAllLeads(dbcfg *config.DbConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		leadsChan := make(chan sqlc.Lead)
		errChan := make(chan error)

		dbcfg.GetDataAsync(leadsChan, errChan, r, "SELECT * FROM LEADS")

		var leads []sqlc.Lead
		var errors []error
		var success int = 0
		var failure int = 0

		for leadsChan != nil || errChan != nil {
			select {
			case lead, ok := <-leadsChan:
				if !ok {
					leadsChan = nil
					continue
				}

				leads = append(leads, lead)
				success++

			case err, ok := <-errChan:
				if !ok {
					errChan = nil
					continue
				}

				errors = append(errors, err)
				failure++
			}
		}

		utils.WriteResponseMiddleware(w, leads, errors, success, failure)
	}
}
