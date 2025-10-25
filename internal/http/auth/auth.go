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

	// here storage is db instance passed from main.go and we can use it to call stuct storage methods like CreateUser

	return func(w http.ResponseWriter, r *http.Request) {
		//response.WriteJson(w, http.StatusOK, map[string]string{"message": "Signup endpoint reached"})
		var details types.Signup
		err := json.NewDecoder(r.Body).Decode(&details)
		if errors.Is(err, io.EOF) {
			fmt.Println("Empty body")
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Empty request body"})
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
			return
		}

		fmt.Println(details)

		id, token, er := storage.CreateUser(&details.Username, &details.Email, &details.Password)
		if er != nil {
			fmt.Println("Error creating user:", er)
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": er.Error()})
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "User created successfully", "user_id": fmt.Sprint(id), "token": token})

	}
}

func Login(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract details ffrom response body
		var details types.Login // details will be of login type struct

		err := json.NewDecoder(r.Body).Decode(&details) // decode the body and put the details in details struct
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Empty request body"})
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
			return
		}
		fmt.Println(details)

		// now we will verify the credentials
		id, token, er := storage.Login(&details.Email, &details.Password)
		if er != nil {
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"error": er.Error()})
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "User logged in Sucessfully", "user_id": fmt.Sprint(id), "token": token})

	}
}
