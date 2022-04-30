package routes

import (
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/controllers"
)

var moodRoutes = func(router *mux.Router) {
	router.HandleFunc("/mood/{user_id}", controllers.getUserMood).Methods("GET")
	router.HandleFunc("/mood/{user_id}/{created_at}", controllers.getMoodByDay).Methods("GET")
	router.HandleFunc("/mood/", controllers.addMoodLog).Methods("POST")
	router.HandleFunc("/mood/{user_id}/{id}", controllers.updateMoodLog).Methods("PUT")
	router.HandleFunc("/mood/{user_id}/{id}", controllers.deleteMoodLog).Methods("DELETE")
}