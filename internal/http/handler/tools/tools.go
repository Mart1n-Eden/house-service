package tools

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request) (t T, err error) {
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}
