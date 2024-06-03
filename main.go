package main

import (
	"first-jwt/configs"
	"first-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// connectDB function from configs packages
	configs.ConnectDB()

	// mux router
	r:=mux.NewRouter()
	
	//router initiation
	router := r.PathPrefix("/api").Subrouter() // This function for making grouping route

	routes.AuthRoutes(router) // sub router from the routes grouping before
	routes.UserRoutes(router) // another subrouting
	routes.BookRoutes(router) // book subrouter

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}