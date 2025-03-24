package main

import (
	"github.com/cocoth/linknet-api/cmd"
	"github.com/cocoth/linknet-api/src/database"
	"github.com/cocoth/linknet-api/src/utils"
)

func init() {
	utils.LoadEnv()
	database.ConnectToDB()
}

// func runServer(e *gin.Engine) {
// 	env, port := os.Getenv("APP_ENV"), os.Getenv("APP_PORT")
// 	var host string

// 	if env == "prod" {
// 		gin.SetMode(gin.ReleaseMode)
// 		host = "0.0.0.0"
// 	} else if env == "dev-docker" {
// 		host = "0.0.0.0"
// 	} else {
// 		host = os.Getenv("APP_HOST")
// 	}
// 	server := host + ":" + port
// 	if err := e.Run(server); err != nil {
// 		utils.Error(err.Error(), "runServer")
// 		os.Exit(1)
// 	}
// }

func main() {
	// db := database.DB

	// rootCmd := cmd.Opt(db)
	cmd.Exec()

	// r := gin.Default()

	// validate := validator.New()

	// v1 := r.Group("/api/v1")

	// userRepo := repo.NewUserRepoImpl(db)

	// userService := services.NewUserServiceImpl(userRepo, validate)
	// authService := services.NewAuthService(userRepo)
	// authMiddleware := middleware.NewUserAuthorization(userService)

	// userCtrl := controllers.NewUserController(userService)
	// authCtrl := controllers.NewAuthController(authService)

	// routes.AuthRoute(authMiddleware, authCtrl, v1)
	// routes.UserRoute(userCtrl, v1)

	// runServer(r)
}
