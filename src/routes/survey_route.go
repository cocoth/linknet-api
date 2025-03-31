package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SurveyRoute(surveyController *controllers.SurveyController, rg *gin.RouterGroup) {
	rg.GET("/surveys", surveyController.GetAllSurvey)
	rg.POST("/surveys", surveyController.CreateSurvey)
	rg.PATCH("/surveys/:id", surveyController.UpdateSurvey)
	rg.DELETE("/surveys/:id", surveyController.DeleteSurvey)
}

func NotificationRoute(notifController *controllers.NotifyController, rg *gin.RouterGroup) {
	rg.GET("/notif", notifController.GetAllNotify)
	rg.POST("/notif", notifController.CreateNotify)
	rg.DELETE("/notif/:id", notifController.DeleteNotify)

	// Websocket
	rg.GET("/ws/notif", controllers.HandleWebsocketConnection)
}
