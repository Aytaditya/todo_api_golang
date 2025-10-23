package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json") // setting content type to json
	w.WriteHeader(status)                              // setting status code
	return json.NewEncoder(w).Encode(data)             // encoding data to json and writing to response writer
}
