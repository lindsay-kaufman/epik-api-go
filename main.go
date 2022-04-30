package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lindsay-kaufman/go-epik-api/app/routes"
)


func main() {
	r := mux.NewRouter()

	routes.authRoutes(r)
	routes.activityRoutes(r)
	routes.goalRoutes(r)
	routes.mealRoutes(r)
	routes.moodRoutes(r)
	routes.sleepRoutes(r)
	routes.toDoRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost/5432", r))
}