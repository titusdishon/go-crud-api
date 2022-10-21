package routes

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/titusdishon/go-docker-mysql/controllers"
	router "github.com/titusdishon/go-docker-mysql/http"
	"os"
)

var (
	httpRouter router.Router               = router.NewMuxRouter()
	controller controllers.IUserController = controllers.NewUserController()
)

var UserRouters = func() {
	err := godotenv.Load()
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err != nil {
		fmt.Println("failed to load env files")
	}
	fmt.Printf("Mux running on port: %s", PORT)
	httpRouter.GET("/", controller.PingMe)
	httpRouter.GET("/users", controller.GetUsers)
	httpRouter.POST("/user/create", controller.CreateUser)
	httpRouter.GET("/user/get-by-id/{userId:[0-9]+}", controller.GetUserById)
	httpRouter.PUT("/user/update/{userId:[0-9]+}", controller.UpdateUser)
	httpRouter.DELETE("/user/delete/{userId:[0-9]+}", controller.DeleteAUser)
	httpRouter.SERVE(PORT)
}
