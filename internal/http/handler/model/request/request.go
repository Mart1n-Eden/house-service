package request

type HouseCreateRequest struct {
	Address   string  `json:"address"`
	Year      int     `json:"year"`
	Developer *string `json:"developer"`
}

type FlatCreateRequest struct {
	HouseId int `json:"house_id"`
	Price   int `json:"price"`
	Rooms   int `json:"rooms"`
}

type FlatUpdateRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type LoginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type SubscriptionRequest struct {
	Email string `json:"email"`
}
