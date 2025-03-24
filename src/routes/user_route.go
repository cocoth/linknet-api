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
	// Add new user
	rg.POST("/user", ctrl.CreateUser)
	// Update Existing user
	rg.PATCH("/user/:id", ctrl.UpdateUser)
	// Delete user
	rg.DELETE("/user/:id", ctrl.DeleteUser)

	// Get all roles
	rg.GET("/user/roles", ctrl.GetAllRole)
	// Add new role
	rg.POST("/user/role", ctrl.CreateRole)
	// Update Existing role
	rg.PATCH("/user/role/:id", ctrl.UpdateRole)
	// Delete role
	rg.DELETE("/user/role/:id", ctrl.DeleteRole)
}

func AuthRoute(authMiddleware *middleware.UserAuthorization, ctrl *controllers.UserAuthController, rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", ctrl.Register)
	// Login
	rg.POST("/login", ctrl.Login)
	// Logout
	rg.POST("/logout", ctrl.Logout)
	// Validate token
	rg.GET("/validate", authMiddleware.Authorize, ctrl.Validate)
}

func FileRoute(ctrl *controllers.FileController, rg *gin.RouterGroup) {
	// Upload file
	rg.POST("/files/upload", ctrl.UploadFile)
	// Get all files
	rg.GET("/files", ctrl.GetAllFileUpload)
	// Download file qurey params (id, fileid, filename, filehash)
	rg.GET("/files/download", ctrl.DownloadFile)
	// Update file
	rg.PATCH("/files/:id", ctrl.UpdateFileUpload)
	// // Get file by qurey params (id, fileid, filename, filehash)
	rg.DELETE("/files", ctrl.DeleteFileUpload)
}
