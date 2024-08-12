package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type HouseResponse struct {
	Id        int       `json:"id"`
	Address   string    `json:"address"`
	Year      int       `json:"year"`
	Developer string    `json:"developer"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

//type FlatResponse struct {
//	Flats []Flat `json:"flats"`
//}

type FlatResponse struct {
	Id      int    `json:"id"`
	HouseId int    `json:"house_id"`
	Price   int    `json:"price"`
	Rooms   int    `json:"rooms"`
	Status  string `json:"status"`
}

type UserIdResponse struct {
	UserId string `json:"user_id"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// TODO: move to handler/tools
func SendResponse(msg any, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

func SendStarus(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
