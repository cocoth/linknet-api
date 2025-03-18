package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/http/middleware"

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
	rg.PATCH("/user/:id", ctrl.Update)
	// Delete user
	rg.DELETE("/user/:id", ctrl.Delete)

	// Get all roles
	rg.GET("/user/roles", ctrl.GetAllRole)
	// Add new role
	rg.POST("/user/role", ctrl.CreateRole)
}

func AuthRoute(authMiddleware *middleware.UserAuthorization, ctrl *controllers.UserAuthController, rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", ctrl.Register)
	// Login
	rg.POST("/login", ctrl.Login)
	// Logout
	rg.POST("/logout", ctrl.Logout)

	rg.GET("/validate", authMiddleware.Authorize, ctrl.Validate)
}
