package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	database, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		panic("could not initialize database")
	}

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)
	DB = database

	createTables()
}

func createTables() {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(usersTable)
	if err != nil {
		str := fmt.Sprintf("could not create users table\nerror: %v", err)
		panic(str)
	}

	todoTable := `
	CREATE TABLE IF NOT EXISTS todos(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT DEFAULT '',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(todoTable)
	if err != nil {
		str := fmt.Sprintf("could not create todos table\nerror: %v", err)
		panic(str)
	}
}
