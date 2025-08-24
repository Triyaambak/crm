package database

import (
	"database/sql"
	"fmt"
	"log"

	helper "github.com/triyaambak/CRM/utils"

	_ "github.com/lib/pq"
)

type DB struct {
	Db *sql.DB
}

func (d *DB) InitDB() {

	var (
		user, password, host, port, dbname string
		err                                error
		db                                 *sql.DB
	)

	h := helper.Helper{}

	user, password, host, port, dbname, err = h.CheckPostgressENV()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	success := false
	for i := 2; i >= 0; i-- {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			fmt.Printf("Connection failed.Retrying %d more times", i)
		} else if err = db.Ping(); err != nil {
			fmt.Printf("Ping failed - Error: %v. Retrying %d more times\n", err, i)
		} else {
			fmt.Println("Connection established successfully")
			success = true
			break
		}
	}

	if !success {
		log.Fatal("Connection to server failed")
	}

	(*d).Db = db

	

}
