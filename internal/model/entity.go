package model

import "time"

type House struct {
	Id        int
	Address   string
	Year      int
	Developer string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type Flat struct {
	Id      int
	HouseId int
	Price   int
	Rooms   int
	Status  string
}

type User struct {
	Id       int
	Email    string
	Password string
	UserType string
}
