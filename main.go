package main

import (
	"HerokuService/Controllers"
	"strings"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"	
)

func main() {
	
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// port := viper.GetString("port")
	
	mainController := Controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/login", mainController.Login()).Methods("POST")
	router.HandleFunc("/profile", mainController.Profile()).Methods("POST")

	fmt.Println("Starting Service....")
	http.ListenAndServe(":8080",router)
}
