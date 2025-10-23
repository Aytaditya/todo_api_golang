package sqlite

import (
	"database/sql"

	"github.com/Aytaditya/todo_api_golang/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	DB *sql.DB
}

// this function will accept a config struct and return a pointer to Sqlite struct and an error
func ConnectDB(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.Storagepath) // open a connection to the sqlite database using the path from config
	if err != nil {
		return nil, err
	}

	// creating user table
	_, er := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)
	if er != nil {
		return nil, er
	}

	// now creating second table (notes table)
	_, er_handle := db.Exec(`CREATE TABLE IF NOT EXISTS notes(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT,
		tag TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id))`)

	if er_handle != nil {
		return nil, er_handle
	}

	return &Sqlite{DB: db}, nil

}
