package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`

	db.Exec(sqlStmt)

	return db
}
