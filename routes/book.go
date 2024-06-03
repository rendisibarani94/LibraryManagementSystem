package routes

import (
	"first-jwt/controllers"
	// "first-jwt/middleware"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router){
	router := r.PathPrefix("/book").Subrouter()

	// router.Use(middleware.Auth)
	router.HandleFunc("/" , controllers.ViewALlBooks).Methods("GET")
	router.HandleFunc("/{id}" , controllers.ViewBooksById).Methods("GET")
	router.HandleFunc("/add" , controllers.AddBook).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/borrow" , controllers.BorrowBook).Methods("POST")
	router.HandleFunc("/borrowed" , controllers.GetAllBorrowedBook).Methods("GET")

}