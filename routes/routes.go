package routes

import (
	"net/http"

	"github.com/titusdishon/go-docker-mysql/controllers"
)

var UserRouters = func() {
	http.HandleFunc("/", controllers.PingMe)
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/user/create", controllers.CreateUser)
}
