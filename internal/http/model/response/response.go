package response

import (
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

type FlatResponse struct {
	Flats []Flat `json:"flats"`
}

type Flat struct {
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

type Error struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
}
