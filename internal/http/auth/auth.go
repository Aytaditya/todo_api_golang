package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Aytaditya/todo_api_golang/internal/response"
	"github.com/Aytaditya/todo_api_golang/internal/storage/sqlite"
	"github.com/Aytaditya/todo_api_golang/internal/types"
)

func Signup(storage *sqlite.Sqlite) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//response.WriteJson(w, http.StatusOK, map[string]string{"message": "Signup endpoint reached"})

		var details types.Signup
		err := json.NewDecoder(r.Body).Decode(&details)
		if errors.Is(err, io.EOF) {
			fmt.Println("Empty body")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Empty request body"}`))
			return
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid JSON format"}`))
			return
		}

		fmt.Println(details)

		id, er := storage.CreateUser(&details.Username, &details.Email, &details.Password)
		if er != nil {
			fmt.Println("Error creating user:", er)
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": er.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		// Set status code
		w.WriteHeader(http.StatusOK)
		// Write JSON response
		w.Write([]byte(`{"message": "User created with ID ` + fmt.Sprint(id) + `"}`))

	}
}
