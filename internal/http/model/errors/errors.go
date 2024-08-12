package errors

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
}

// TODO: mode to handler/tools
func ResponseError(msg string, w http.ResponseWriter, code int) {
	// TODO: other fields
	e := Error{
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}
