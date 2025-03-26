package routes

import (
	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SurveyRoute(surveyController *controllers.SurveyController, rg *gin.RouterGroup) {
	rg.GET("/surveys", surveyController.GetAllSurvey)
	// rg.GET("/survey/:id", surveyController.GetSurveyByID)
	rg.POST("/surveys", surveyController.CreateSurvey)
	// rg.PUT("/survey/:id", surveyController.UpdateSurvey)
	// rg.DELETE("/survey/:id", surveyController.DeleteSurvey)
}
