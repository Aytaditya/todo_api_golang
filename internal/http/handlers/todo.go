package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aytaditya/todo_api_golang/internal/middleware/jwt"
	"github.com/Aytaditya/todo_api_golang/internal/response"
	"github.com/Aytaditya/todo_api_golang/internal/storage/sqlite"
	"github.com/Aytaditya/todo_api_golang/internal/types"
)

func CreateTodo(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// first we will look into the jwt token from the authorization header (currently doing without JWT Middleware but recmoneded to use middleware for better code structure)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2) // this will split the header into two parts "Bearer" and the token
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
			return
		}
		tokenString := parts[1] // this is the actual token part

		// now we will validate the token using our jwt middleware
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			return
		}

		userId := claims.ID // extracting user id from the claims
		fmt.Println("Authenticated user ID:", userId)

		// now we will extract the json body
		var tododata types.CreateTodo
		er := json.NewDecoder(r.Body).Decode(&tododata)
		if er != nil {
			response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
			return
		}

		todoId, ok := storage.CreatingTodo(&userId, &tododata.Title, &tododata.Content, &tododata.Tag)
		if ok != nil {
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"error": ok.Error()})
			return
		}
		response.WriteJson(w, http.StatusOK, map[string]interface{}{"message": "Todo created successfully", "todo_id": todoId})

	}
}
