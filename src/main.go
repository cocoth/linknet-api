package main

import (
	"os"

	"github.com/cocoth/linknet-api/config"
	"github.com/cocoth/linknet-api/config/options"
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/http/middleware"
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

func runServer(e *gin.Engine) {
	env, port := os.Getenv("APP_ENV"), os.Getenv("APP_PORT")
	var host string

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		host = "0.0.0.0"
	} else if env == "dev-docker" {
		host = "0.0.0.0"
	} else {
		host = os.Getenv("APP_HOST")
	}
	server := host + ":" + port
	if err := e.Run(server); err != nil {
		panic(err)
	}
}

func main() {
	db := config.DB

	options.Opt(db)
	r := gin.Default()

	validate := validator.New()

	v1 := r.Group("/api/v1")

	userRepo := repo.NewUserRepoImpl(db)

	userService := services.NewUserServiceImpl(userRepo, validate)
	authService := services.NewAuthService(userRepo)
	authMiddleware := middleware.NewUserAuthorization(userService)

	userCtrl := controllers.NewUserController(userService)
	authCtrl := controllers.NewAuthController(authService)

	routes.AuthRoute(authMiddleware, authCtrl, v1)
	routes.UserRoute(userCtrl, v1)

	runServer(r)
}
