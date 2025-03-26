package controllers

import (
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
	var surveys []response.SurveyResponse
	var err error

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

	if qID != "" {
		survey, err := s.surveyService.GetSurveyByID(qID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qTitle != "" {
		surveys, err = s.surveyService.GetSurveysByTitle(qTitle)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qFormNumber != "" {
		survey, err := s.surveyService.GetSurveyByFormNumber(qFormNumber)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qQuestorName != "" {
		surveys, err = s.surveyService.GetSurveysByQuestorName(qQuestorName)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qFAT != "" {
		survey, err := s.surveyService.GetSurveyByFAT(qFAT)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qCustomerName != "" {
		surveys, err = s.surveyService.GetSurveysByCustomerName(qCustomerName)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qAddress != "" {
		surveys, err = s.surveyService.GetSurveysByAddress(qAddress)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qNodeFDT != "" {
		survey, err := s.surveyService.GetSurveyByNodeFDT(qNodeFDT)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qSurveyDate != "" {
		survey, err := s.surveyService.GetSurveyBySurveyDate(qSurveyDate)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qSurveyorID != "" {
		survey, err := s.surveyService.GetSurveyBySurveyorID(qSurveyorID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else if qImageID != "" {
		survey, err := s.surveyService.GetSurveyByImageID(qImageID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, survey)
		return
	} else {
		surveys, err = s.surveyService.GetAllSurvey()
	}

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, surveys)
}

func (s *SurveyController) CreateSurvey(c *gin.Context) {
	var surveyReq request.SurveyRequest

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	_, _, errToken := s.userService.CheckToken(token)
	if errToken != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errToken.Error())
		return
	}

	if err := c.ShouldBindJSON(&surveyReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	surveyRes, err := s.surveyService.CreateSurvey(surveyReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusCreated, surveyRes)
}
