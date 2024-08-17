package response

import (
	"time"

	"house-service/internal/domain"
)

type HouseResponse struct {
	Id        int       `json:"id"`
	Address   string    `json:"address"`
	Year      int       `json:"year"`
	Developer string    `json:"developer"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type ListFlatsResponse struct {
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

func CreateFlatResponse(f *domain.Flat) Flat {
	return Flat{
		Id:      f.Id,
		HouseId: f.HouseId,
		Price:   f.Price,
		Rooms:   f.Rooms,
		Status:  f.Status,
	}
}

func CreateHouseResponse(h *domain.House) HouseResponse {
	return HouseResponse{
		Id:        h.Id,
		Address:   h.Address,
		Year:      h.Year,
		Developer: h.Developer,
		CreatedAt: h.CreatedAt,
		UpdateAt:  h.UpdatedAt,
	}
}

func CreateListFlatsResponse(h []domain.Flat) ListFlatsResponse {
	flats := make([]Flat, len(h))
	for i, flat := range h {
		flats[i] = Flat{
			Id:      flat.Id,
			HouseId: flat.HouseId,
			Price:   flat.Price,
			Rooms:   flat.Rooms,
			Status:  flat.Status,
		}
	}
	return ListFlatsResponse{Flats: flats}
}

func CreateUserIdResponse(id string) UserIdResponse {
	return UserIdResponse{UserId: id}
}

func CreateTokenResponse(token string) TokenResponse {
	return TokenResponse{Token: token}
}
