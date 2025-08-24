package utils

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	sqlc "github.com/triyaambak/CRM/internal/sqlc_db"
)

type Helper struct{}

func (h *Helper) CheckInputData(name, phone string) (lead sqlc.CreateLeadParams, err error) {

	lead = sqlc.CreateLeadParams{}

	if name == "" {
		err = errors.New("Field name cannot be null")
		return sqlc.CreateLeadParams{}, err
	}

	lead.ID = uuid.New()
	lead.Name = name
	lead.Phone = sql.NullString{String: phone, Valid: phone != ""}
	lead.CreatedAt = time.Now()

	return lead, nil
}

func (h *Helper) CheckPostgressENV() (user, password, host, port, dbname string, err error) {
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	dbname = os.Getenv("POSTGRES_DB")

	err = nil

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		err = errors.New("Invalid Postgres parameters , please check ENV variables")
		return "", "", "", "", "", err
	}

	return
}
