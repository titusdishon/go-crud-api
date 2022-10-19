package routes

import (
	"github.com/gorilla/mux"
	"github.com/titusdishon/go-docker-mysql/controllers"
)

var UserRouters = func(router *mux.Router) {
	router.HandleFunc("/", controllers.PingMe).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/delete/{userId}", controllers.DeleteAUser).Methods("DELETE")
}
