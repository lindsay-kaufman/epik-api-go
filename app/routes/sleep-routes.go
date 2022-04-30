package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var sleepRoutes = func(router *mux.Router) {
	router.HandleFunc("/sleep/{user_id}", controllers.getUserSleep).Methods("GET")
	router.HandleFunc("/sleep/", controllers.addSleepRecord).Methods("POST")
	router.HandleFunc("/sleep/{user_id}/{id}", controllers.updateSleepRecord).Methods("PUT")
	router.HandleFunc("/sleep/{user_id}/{id}", controllers.deleteSleepRecord).Methods("DELETE")
}