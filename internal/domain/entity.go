package domain

import "time"

type House struct {
	Id        int
	Address   string
	Year      int
	Developer string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Flat struct {
	Id      int
	HouseId int
	Price   int
	Rooms   int
	Status  string
}

type User struct {
	Id       string
	Email    string
	Password string
	UserType string
}
