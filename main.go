package main

import (
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
	// port := viper.GetString("port")
	

	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")

	fmt.Println("Starting RESTFUL....")
	http.ListenAndServe(":8080",router)
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
