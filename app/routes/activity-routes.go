package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var activityRoutes = func(router *mux.Router) {
	router.HandleFunc("/activity/{user_id}/{id}", controllers.getUserActivityById).Methods("GET")
	router.HandleFunc("/activity/{user_id}", controllers.getUserActivities).Methods("GET")
	router.HandleFunc("/activity/", controllers.addActivity).Methods("POST")
	router.HandleFunc("/activity/{user_id}/{id}", controllers.updateActivity).Methods("PUT")
	router.HandleFunc("users/{user_id}/{id}", controllers.deleteActivity).Methods("DELETE")
}