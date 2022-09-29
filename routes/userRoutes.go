package routes

import (
	"github.com/RahulG606/go-API/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router){
	router.HandleFunc("/user",controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/user/{userId}",controllers.UpdateAUser()).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
    router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
	
	
}