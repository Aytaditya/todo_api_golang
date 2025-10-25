package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/Aytaditya/todo_api_golang/internal/config"
	"github.com/Aytaditya/todo_api_golang/internal/http/auth"
	"github.com/Aytaditya/todo_api_golang/internal/storage/sqlite"
)

func main() {
	fmt.Println("This is the main package for the todo_api command.")
	cfg := config.MustLoad() // loading all the configurations from the config file (contains detail like port address, database path)

	// DB CONNECTION HERE (it is returning a db instance)
	storage, er := sqlite.ConnectDB(cfg)
	if er != nil {
		log.Fatal("failed to connect to database:", er.Error())
	}

	slog.Info("Connected to the database successfully")

	router := http.NewServeMux() // Create a new HTTP request multiplexer (router)

	// route 1 :
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo API"))
	})

	// AUTH ROUTES
	router.HandleFunc("POST /api/signup", auth.Signup(storage))
	router.HandleFunc("POST /api/login", auth.Login(storage))

	// NOTES ROUTES
	// creating notes route
	// view all notes of a user
	// update a note
	// delete a note

	// http.Server is a struct that represents an HTTP server. here we are creating an instance of http.Server with the specified address and handler (router).
	// but we can also create a server using http.ListenAndServe directly without creating an instance of http.Server.
	// but this gives us more control
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}

	fmt.Println("Server Running at:", cfg.HttpServer.Address)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}

}
