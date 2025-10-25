package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Aytaditya/todo_api_golang/internal/config"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

	// before adding password to db. hash it using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %v", err)
	}

	stmt, err := s.DB.Prepare("INSERT INTO users(username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(*username, *email, string(hashedPassword))
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

func (s *Sqlite) Login(email *string, password *string) (int64, error) {
	if email == nil || password == nil {
		return 0, fmt.Errorf("email and password must not be nil")
	}

	// Fetch user by email
	row := s.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", *email)

	var id int64
	var dbPassword string

	err := row.Scan(&id, &dbPassword)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("no user found with the given email")
	}
	if err != nil {
		return 0, err
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(*password))
	if err != nil {
		return 0, fmt.Errorf("wrong password entered")
	}

	return id, nil
}
