package sqlite

import (
	"database/sql"
	"fmt"

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

func (s *Sqlite) CreateUser(username *string, email *string, password *string) (int64, error) {

	if username == nil || email == nil || password == nil {
		return 0, fmt.Errorf("username, email, and password must not be nil")
	}
	stmt, err := s.DB.Prepare("INSERT INTO users(username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(*username, *email, *password)
	stmt.Close()
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
