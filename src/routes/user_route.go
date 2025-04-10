package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/http/middlewares"

	"github.com/gin-gonic/gin"
)

// func UserRoute(rg *gin.RouterGroup) {
func UserRoute(authMiddleware *middlewares.UserAuthorization, ctrl *controllers.UserController, rg *gin.RouterGroup) {
	// Get all users
	rg.GET("/users", authMiddleware.Authorize, ctrl.GetAll)
	// Add new user
	rg.POST("/user", authMiddleware.Authorize, ctrl.CreateUser)
	// Update Existing user
	rg.PATCH("/user/:id", authMiddleware.Authorize, ctrl.UpdateUser)
	// Delete user
	rg.DELETE("/user/:id", authMiddleware.Authorize, ctrl.DeleteUser)

	// Get all roles
	rg.GET("/user/roles", authMiddleware.Authorize, ctrl.GetAllRole)
	// Add new role
	rg.POST("/user/role", authMiddleware.Authorize, ctrl.CreateRole)
	// Update Existing role
	rg.PATCH("/user/role/:id", authMiddleware.Authorize, ctrl.UpdateRole)
	// Delete role
	rg.DELETE("/user/role/:id", authMiddleware.Authorize, ctrl.DeleteRole)
}

func AuthRoute(authMiddleware *middlewares.UserAuthorization, ctrl *controllers.UserAuthController, rg *gin.RouterGroup) {
	// Register new user
	rg.POST("/register", ctrl.Register)
	// Login
	rg.POST("/login", ctrl.Login)
	// Logout
	rg.POST("/logout", ctrl.Logout)
	// Validate token
	rg.GET("/validate", authMiddleware.Authorize, ctrl.Validate)
	rg.GET("/admin-validate", authMiddleware.Authorize, ctrl.ValidateAdmin)
	// Validate Admin
	rg.POST("/admin-login", ctrl.CheckIsAdmin)
}

func FileRoute(authMiddleware *middlewares.UserAuthorization, ctrl *controllers.FileController, rg *gin.RouterGroup) {
	// Upload file
	rg.POST("/files/upload", authMiddleware.Authorize, ctrl.UploadFile)
	// Get all files
	rg.GET("/files", authMiddleware.Authorize, ctrl.GetAllFileUpload)
	// Download file qurey params (id, fileid, filename, filehash)
	rg.GET("/files/download", authMiddleware.Authorize, ctrl.DownloadFile)
	// Update file
	rg.PATCH("/files/:id", authMiddleware.Authorize, ctrl.UpdateFileUpload)
	// // Get file by qurey params (id, fileid, filename, filehash)
	rg.DELETE("/files", authMiddleware.Authorize, ctrl.DeleteFileUpload)

	rg.POST("/files/request", authMiddleware.Authorize, ctrl.RequestAccess)
	rg.POST("/files/approve", authMiddleware.Authorize, ctrl.ApproveAccess)
	rg.POST("/files/reject", authMiddleware.Authorize, ctrl.RejectAccess)
}
