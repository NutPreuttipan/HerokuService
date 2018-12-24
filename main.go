package main

import (
	"HerokuService/models"
	"strings"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"	
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDescription struct {
	Name string `json:"name"`
	LastName string `json:"lastName"`
	Age string `json:"age"`
}

type CoreResponse struct {
	ID int `json:"id"`
	Description string `json:"desc"`
}

type UserResponse struct {
	ApiResponse CoreResponse `json:"apiResponse"`
	Data []UserDescription `json:"data"`
}

func main() {
	
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	port := viper.GetString("port")
	

	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/profile", profile).Methods("POST")

	fmt.Println("Starting RESTFUL....")
	http.ListenAndServe(":"+port,router)
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type","application/json")

	var auth UserAuth
	var response UserResponse
	var user UserDescription

	json.NewDecoder(r.Body).Decode(&auth)

	if auth.Username == "admin" && auth.Password == "admin123" {
		response.ApiResponse.ID = 0
		response.ApiResponse.Description = "Success"
		
		user.Name = "Preuttipan"
		user.LastName = "Janpen"
		user.Age = "26"

		response.Data = append(response.Data, user)

		json.NewEncoder(w).Encode(response)
	} else {
		response.ApiResponse.ID = -1
		response.ApiResponse.Description = "Failed"

		json.NewEncoder(w).Encode(response)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")

	var user UserDescription
	var userDetail models.UserModel
	var homeAddress []models.AddressDetail
	var workAddress []models.AddressDetail
	var response models.UserModelResponse

	homeAddress = append(homeAddress,
		models.AddressDetail{ID:1,Address1:"797",District:"เมือง",Province:"นนทบุรี",Zipcode:"11000"},
		models.AddressDetail{ID:2,Address1:"5/2520",District:"ท่าทราย",Province:"นนทบุรี",Zipcode:"12000"})

	workAddress = append(workAddress,
		models.AddressDetail{ID:3,Address1:"219/1",District:"วิเชียร์บุรี",Province:"เพชรบูรณ์",Zipcode:"67180"})

	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "Preuttipan" {

		response.ApiResponse.ID = 0
		response.ApiResponse.Description = "Success"

		userDetail.ID = 1
		userDetail.Name = "Preuttipan"
		userDetail.Lname = "Janpen"
		userDetail.Address.Home = homeAddress
		userDetail.Address.Work = workAddress

		response.Data = append(response.Data,userDetail)

		json.NewEncoder(w).Encode(response)
	} else {
		response.ApiResponse.ID = -1
		response.ApiResponse.Description = "Failed"

		json.NewEncoder(w).Encode(response)
	}
}
