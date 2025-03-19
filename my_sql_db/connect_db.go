package my_sql_db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // <-- Ensure this import is present
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/info")
	if err != nil {
		return nil, err
	}
	return db, nil
}
