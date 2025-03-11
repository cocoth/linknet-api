package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"

	"github.com/gin-gonic/gin"
)

// func UserRoute(rg *gin.RouterGroup) {
func UserRoute(ctrl *controllers.UserController, rg *gin.RouterGroup) {
	// Get all users
	rg.GET("/users", ctrl.GetAll)
	// Get user by id
	rg.GET("/user/:id", ctrl.GetById)
	// Add new user
	rg.POST("/user", ctrl.Create)
	// Update Existing user
	rg.PUT("/user/:id", ctrl.Update)
	// Dele user
	rg.DELETE("/user/:id", ctrl.Delete)
}

func AuthRoute(ctrl *controllers.UserAuthController, rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", ctrl.Register)
	// Login
	rg.POST("/login", ctrl.Login)
	// Logout
	// rg.POST("/logout", ctrl.Logout)
}
