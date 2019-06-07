package Controllers

import (
	"encoding/json"
	"net/http"
	models "HerokuService/Models"
)

type Controller struct{}

func (c Controller) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
	
		var auth models.UserAuth
		var response models.LoginResponse
		var user models.UserDescription
	
		json.NewDecoder(r.Body).Decode(&auth)
	
		if auth.Username == "admin" && auth.Password == "admin123" {
			response.ApiResponse.ID = 0
			response.ApiResponse.Description = "Success"
			 
			user.Name = "Nuttarikan"
			user.LastName = "Pattanaporn"
			user.Age = "23"
	
			response.Data = append(response.Data, user)
	
			json.NewEncoder(w).Encode(response)
		} else {
			response.ApiResponse.ID = -1
			response.ApiResponse.Description = "Failed"
	
			json.NewEncoder(w).Encode(response)
		}
	}	
}

func (c Controller) Profile() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type","application/json")
	
		var user models.UserDescription
		var userDetail models.UserModel
		var homeAddress []models.AddressDetail
		var workAddress []models.AddressDetail
		var response models.ProfileResponse
	
		homeAddress = append(homeAddress,
			models.AddressDetail{ID:1,Address1:"797",District:"เมือง",Province:"นนทบุรี",Zipcode:"11000"},
			models.AddressDetail{ID:2,Address1:"5/2520",District:"ท่าทราย",Province:"นนทบุรี",Zipcode:"12000"})
	
		workAddress = append(workAddress,
			models.AddressDetail{ID:3,Address1:"219/1",District:"วิเชียร์บุรี",Province:"เพชรบูรณ์",Zipcode:"67180"})
	
		json.NewDecoder(r.Body).Decode(&user)
	
		if user.Name == "Nuttarikan" {
	
			response.ApiResponse.ID = 0
			response.ApiResponse.Description = "Success"
	
			userDetail.ID = 1
			userDetail.Name = "Nuttarikan"
			userDetail.Lname = "Pattanaporn"
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
}

