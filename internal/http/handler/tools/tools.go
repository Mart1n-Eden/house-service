package tools

import (
	"encoding/json"
	"net/http"

	"house-service/internal/http/model/response"
)

func Decode[T any](r *http.Request) (t T, err error) {
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func SendResponse(w http.ResponseWriter, msg any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

func SendInternalError(w http.ResponseWriter, msg string, code int) {
	e := response.Error{
		Message: msg,
		// TODO: RequestId
		Code: code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}

func SendStarus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
