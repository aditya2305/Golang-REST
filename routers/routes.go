package routes

import (
	controller "github.com/aditya2305/REST_API/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Routes consist of a path and a handler function
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movies", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", controller.DeleteAMovie).Methods("DELETE")

	return router
}
