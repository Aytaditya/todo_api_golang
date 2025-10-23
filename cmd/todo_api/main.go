package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aytaditya/todo_api_golang/internal/config"
)

func main() {
	fmt.Println("This is the main package for the todo_api command.")
	cfg := config.MustLoad() // loading all the configurations from the config file (contains detail like port address, database path)

	fmt.Println("server at address:", cfg.Address)

	router := http.NewServeMux() // Create a new HTTP request multiplexer (router)

	// route 1 :
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Todo API"))
	})

	// http.Server is a struct that represents an HTTP server. here we are creating an instance of http.Server with the specified address and handler (router).
	// but we can also create a server using http.ListenAndServe directly without creating an instance of http.Server.
	// but this gives us more control
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}

	err := server.ListenAndServe()

	fmt.Println("Server Running at:", cfg.HttpServer.Address)

	if err != nil {
		log.Fatal(err.Error())
	}

}
