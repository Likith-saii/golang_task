package my_sql_db

import "database/sql"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/info")
	if err != nil {
		return nil, err
	}
	return db, nil
}
