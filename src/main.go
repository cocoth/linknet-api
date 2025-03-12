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

	env, port := os.Getenv("APP_ENV"), os.Getenv("APP_PORT")
	var host string

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		host = "0.0.0.0"
	} else if env == "dev-docker" {
		host = "0.0.0.0"
	} else {
		host = "localhost"
	}

	server := host + ":" + port
	if err := r.Run(server); err != nil {
		panic(err)
	}
}
