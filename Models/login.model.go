package Models

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDescription struct {
	Name string `json:"name"`
	LastName string `json:"lastName"`
	Age string `json:"age"`
}

type LoginResponse struct {
	ApiResponse CoreResponse `json:"apiResponse"`
	Data []UserDescription `json:"data"`
}