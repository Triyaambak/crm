package config

import (
	"net/http"

	"github.com/triyaambak/CRM/database"
	sqlc "github.com/triyaambak/CRM/internal/sqlc_db"
)

type DbConfig struct {
	db    *database.DB
	Query *sqlc.Queries
}

func (d *DbConfig) Close() {
	d.db.Db.Close()
}

// This function is used to get the data asynchronously
// SQLC generated functions do not give async functions
// So we need to write our own go functions which implement that
// we dont need to specify * (pointers) in go , since channel by default is passed by reference
func (d *DbConfig) GetDataAsync(leadsChan chan sqlc.Lead, errChan chan error, r *http.Request, query string) {

	rows, err := d.db.Db.QueryContext(r.Context(), "SELECT * FROM LEADS")
	if err != nil {
		errChan <- err
		close(leadsChan)
		close(errChan)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var lead sqlc.Lead
		if err := rows.Scan(&lead.ID, &lead.Name, &lead.Phone, &lead.CreatedAt); err != nil {
			errChan <- err
			break
		}

		leadsChan <- lead
	}

	close(leadsChan)
	close(errChan)

}

func NewDBConfig() *DbConfig {
	var db database.DB
	db.InitDB()

	return &DbConfig{
		db:    &db,
		Query: sqlc.New(db.Db),
	}
}
