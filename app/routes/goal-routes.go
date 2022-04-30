package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var goalRoutes = func(router *mux.Router) {
	router.HandleFunc("/goals/{user_id}", controllers.getUserGoals).Methods("GET")
	router.HandleFunc("/goals/{user_id}/{id}", controllers.getGoalById).Methods("GET")
	router.HandleFunc("/goals/", controllers.addGoal).Methods("POST")
	router.HandleFunc("/goals/{user_id}/{id}", controllers.updateGoal).Methods("PUT")
	router.HandleFunc("/goals/{user_id}/{id}", controllers.deleteGoal).Methods("DELETE")
}