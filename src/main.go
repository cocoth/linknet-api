package main

import (
	"os"

	"github.com/cocoth/linknet-api/config"
	"github.com/cocoth/linknet-api/config/options"
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/routes"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
	config.ConnectToDB()
}

func main() {
	options.Opt()
	r := gin.Default()

	db := config.DB
	validate := validator.New()

	v1 := r.Group("/api/v1")

	userRepo := repo.NewUserRepoImpl(db)

	userService := services.NewUserServiceImpl(userRepo, validate)
	authService := services.NewAuthService(userRepo)

	userCtrl := controllers.NewUserController(userService)
	authCtrl := controllers.NewAuthController(authService)

	routes.AuthRoute(authCtrl, v1)
	routes.UserRoute(userCtrl, v1)

	host, port := os.Getenv("APP_HOST"), os.Getenv("APP_PORT")
	server := host + ":" + port

	r.Run(server)
}
