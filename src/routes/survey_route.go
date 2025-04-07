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

	rg.GET("/surveys/view/:id", authMiddleware.Authorize, surveyController.ViewSurveyAndReportsByID)
	rg.GET("/surveys/download/:id", authMiddleware.Authorize, surveyController.DownloadSurveyAndReportsByID)
}

func SurveyReportRoute(authMiddleware *middlewares.UserAuthorization, surveyReportController *controllers.SurveyReportController, rg *gin.RouterGroup) {
	rg.GET("/reports", surveyReportController.GetAllReport)
	rg.POST("/reports", authMiddleware.Authorize, surveyReportController.CreateSurveyReport)
	rg.PATCH("/reports/:id", authMiddleware.Authorize, surveyReportController.UpdateSurveyReport)
	rg.DELETE("/reports/:id", authMiddleware.Authorize, surveyReportController.DeleteSurveyReport)
}

func NotificationRoute(authMiddleware *middlewares.UserAuthorization, notifController *controllers.NotifyController, rg *gin.RouterGroup) {
	rg.GET("/notif", authMiddleware.Authorize, notifController.GetAllNotify)
	rg.POST("/notif", authMiddleware.Authorize, notifController.CreateNotify)
	rg.DELETE("/notif/:id", authMiddleware.Authorize, notifController.DeleteNotify)

	// Websocket
	rg.GET("/ws/notif", authMiddleware.Authorize, controllers.HandleWebsocketConnection)
}
