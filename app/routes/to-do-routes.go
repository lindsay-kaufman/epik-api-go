package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var toDoRoutes = func(router *mux.Router) {
	router.HandleFunc("/todo/{user_id}", controllers.getToDoList).Methods("GET")
	router.HandleFunc("/todo/", controllers.addToDoListItem).Methods("POST")
	router.HandleFunc("/todo/{user_id}/{id}", controllers.updateToDoListItem).Methods("PUT")
	router.HandleFunc("/todo/{user_id}/{id}", controllers.deleteToDoListItem).Methods("DELETE")
}