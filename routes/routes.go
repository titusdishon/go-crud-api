package routes

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/titusdishon/go-docker-mysql/controllers"
	router "github.com/titusdishon/go-docker-mysql/http"
	"github.com/titusdishon/go-docker-mysql/repositories"
	"github.com/titusdishon/go-docker-mysql/services"
)

var (
	repo       repositories.UserRepository = repositories.NewMysqlRepository()
	service    services.UserService        = services.NewUserService(repo)
	httpRouter router.Router               = router.NewMuxRouter()
	controller controllers.IUserController = controllers.NewUserController(service)
)

var UserRouters = func() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env files")
	}
	PORT := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	httpRouter.GET("/", controller.PingMe)
	httpRouter.GET("/users", controller.GetUsers)
	httpRouter.POST("/user/create", controller.CreateUser)
	httpRouter.POST("/user/sign-in", controller.SignIn)
	httpRouter.GET("/user/get-by-id/{userId:[0-9]+}", controller.GetUserById)
	httpRouter.PUT("/user/update/{userId:[0-9]+}", controller.UpdateUser)
	httpRouter.DELETE("/user/delete/{userId:[0-9]+}", controller.DeleteAUser)
	httpRouter.SERVE(PORT)
}
