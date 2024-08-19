package tools

import (
	"encoding/json"
	"net/http"

	"house-service/internal/http/handler/model/response"
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

func SendClientError(w http.ResponseWriter, msg string, code int) {
	err := response.ErrorClient{
		Error: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func SendInternalError(w http.ResponseWriter, msg string, requestId string, code int) {
	e := response.ErrorInternal{
		Message:   msg,
		RequestId: requestId,
		Code:      code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}

func SendStarus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
