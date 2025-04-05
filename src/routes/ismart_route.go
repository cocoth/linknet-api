package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func ISmartRoute(authMiddleware *middlewares.UserAuthorization, ctrl *controllers.ISmartControler, rg *gin.RouterGroup) {
	// Get all ISmart
	rg.GET("/ismart", authMiddleware.Authorize, ctrl.GetAllISmart)
	// Add new ISmart
	rg.POST("/ismart", authMiddleware.Authorize, ctrl.CreateISmart)
	// Update Existing ISmart
	rg.PATCH("/ismart/:id", authMiddleware.Authorize, ctrl.UpdateISmart)
	// Delete ISmart
	rg.DELETE("/ismart/:id", authMiddleware.Authorize, ctrl.DeleteISmart)
}
