package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var authRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/", controllers.signIn).Methods("POST")
	router.HandleFunc("/users/", controllers.signUp).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.deleteUser).Methods("DELETE")
}