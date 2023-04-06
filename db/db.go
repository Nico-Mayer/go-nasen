package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nico-mayer/go-api/config"
)

var DB *sql.DB

func Connect() error {
	connStr := "user=" + config.PGUSER + " password=" + config.PGPASSWORD + " dbname=" + config.PGDATABASE + " host=" + config.PGHOST + " sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr) // Replace "postgres" with the appropriate driver for your database
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
