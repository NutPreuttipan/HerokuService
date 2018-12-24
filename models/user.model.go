package models

type UserModel struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Lname string `json:"lastname"`
	Address Address `json:"address"`
}

type Address struct {
	Home []AddressDetail `json:"home"`
	Work []AddressDetail `json:"work"`
}

type AddressDetail struct {
	ID int `json:"id"`
	Address1 string `json:"address1"`
	District string `json:"district"`
	Province string `json:"province"`
	Zipcode string `json:"zipcode"`
}

type CoreResponse struct {
	ID int `json:"id"`
	Description string `json:"desc"`
}

type UserModelResponse struct {
	ApiResponse CoreResponse `json:"apiResponse"`
	Data []UserModel `json:"data"`
}