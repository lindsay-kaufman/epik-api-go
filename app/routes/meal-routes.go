package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var mealRoutes = func(router *mux.Router) {
	router.HandleFunc("/meals/{user_id}", controllers.getUserMeals).Methods("GET")
	router.HandleFunc("/meals/", controllers.addNewMeal).Methods("POST")
	router.HandleFunc("/meals/{user_id}/{id}", controllers.updateMeal).Methods("PUT")
	router.HandleFunc("/meals/{user_id}/{id}", controllers.deleteMeal).Methods("DELETE")
}