package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_basic_sql")
	db, err := sql.Open("mysql", "root:H3ru@mysql@tcp(127.0.0.1:3306)/golang_basic_sql")

	if err != nil {
		return nil, err
	}

	return db, nil
}
