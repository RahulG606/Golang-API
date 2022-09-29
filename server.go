package main

import (
	"log"
	"net/http"

	"github.com/RahulG606/go-API/configs"
	"github.com/RahulG606/go-API/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// connecting database

	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":8080",router))


}

