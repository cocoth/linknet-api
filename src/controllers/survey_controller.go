package controllers

import (
	"fmt"
	"net/http"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type SurveyController struct {
	surveyService services.SurveyService
	userService   services.UserService
}

func NewSurveyController(surveyService services.SurveyService, userService services.UserService) *SurveyController {
	return &SurveyController{
		surveyService: surveyService,
		userService:   userService,
	}
}

func (s *SurveyController) GetAllSurvey(c *gin.Context) {

	qID := c.Query("id")
	qTitle := c.Query("title")
	qFormNumber := c.Query("form_number")
	qQuestorName := c.Query("questor_name")
	qFAT := c.Query("fat")
	qCustomerName := c.Query("customer_name")
	qAddress := c.Query("address")
	qNodeFDT := c.Query("node_fdt")
	qSurveyDate := c.Query("survey_date")
	qSurveyorID := c.Query("surveyor_id")
	qImageID := c.Query("image_id")

	filters := map[string]interface{}{}

	if qID != "" {
		filters["id"] = qID
	}
	if qTitle != "" {
		filters["title"] = qTitle
	}
	if qFormNumber != "" {
		filters["form_number"] = qFormNumber
	}
	if qQuestorName != "" {
		filters["questor_name"] = qQuestorName
	}
	if qFAT != "" {
		filters["fat"] = qFAT
	}
	if qCustomerName != "" {
		filters["customer_name"] = qCustomerName
	}
	if qAddress != "" {
		filters["address"] = qAddress
	}
	if qNodeFDT != "" {
		filters["node_fdt"] = qNodeFDT
	}
	if qSurveyDate != "" {
		filters["survey_date"] = qSurveyDate
	}
	if qSurveyorID != "" {
		filters["surveyor_id"] = qSurveyorID
	}
	if qImageID != "" {
		filters["image_id"] = qImageID
	}

	surveys, err := s.surveyService.GetSurveysWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(surveys) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No Surveys found with that given filters")
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, surveys)
}

func (s *SurveyController) CreateSurvey(c *gin.Context) {
	var surveyReq request.SurveyRequest

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create survey!")
		return
	}

	if err := c.ShouldBindJSON(&surveyReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	for _, surveyor := range surveyReq.Surveyors {
		_, err := s.userService.GetUserById(surveyor.SurveyorID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, fmt.Sprintf("Surveyor with ID %s does not exist", surveyor.SurveyorID))
			return
		}
	}
	surveyRes, err := s.surveyService.CreateSurvey(surveyReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusCreated, surveyRes)
}

func (s *SurveyController) UpdateSurvey(c *gin.Context) {
	var surveyReq request.UpdateSurveyRequest

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can update survey!")
		return
	}

	if err := c.ShouldBindJSON(&surveyReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	surveyID := c.Param("id")
	surveyRes, err := s.surveyService.UpdateSurvey(surveyID, surveyReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, surveyRes)
}

func (s *SurveyController) DeleteSurvey(c *gin.Context) {

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can delete survey!")
		return
	}

	surveyID := c.Param("id")

	surveyRes, err := s.surveyService.DeleteSurvey(surveyID)
	if err != nil {
		if err.Error() == "record not found" {
			helper.RespondWithError(c, http.StatusNotFound, err.Error())
		} else {
			helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, surveyRes)
}
