package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func SurveyRoute(authMiddleware *middlewares.UserAuthorization, surveyController *controllers.SurveyController, rg *gin.RouterGroup) {
	rg.GET("/surveys", surveyController.GetAllSurvey)
	rg.POST("/surveys", authMiddleware.Authorize, surveyController.CreateSurvey)
	rg.PATCH("/surveys/:id", authMiddleware.Authorize, surveyController.UpdateSurvey)
	rg.DELETE("/surveys/:id", authMiddleware.Authorize, surveyController.DeleteSurvey)
}

func SurveyReportRoute(authMiddleware *middlewares.UserAuthorization, surveyReportController *controllers.SurveyReportController, rg *gin.RouterGroup) {
	rg.GET("/reports", surveyReportController.GetAllReport)
	rg.POST("/reports", authMiddleware.Authorize, surveyReportController.CreateSurveyReport)
	rg.PATCH("/reports/:id", authMiddleware.Authorize, surveyReportController.UpdateSurveyReport)
	rg.DELETE("/reports/:id", authMiddleware.Authorize, surveyReportController.DeleteSurveyReport)
}

func NotificationRoute(notifController *controllers.NotifyController, rg *gin.RouterGroup) {
	rg.GET("/notif", notifController.GetAllNotify)
	rg.POST("/notif", notifController.CreateNotify)
	rg.DELETE("/notif/:id", notifController.DeleteNotify)

	// Websocket
	rg.GET("/ws/notif", controllers.HandleWebsocketConnection)
}
