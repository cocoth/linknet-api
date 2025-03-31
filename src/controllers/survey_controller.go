package controllers

import (
	"fmt"
	"net/http"
	"time"

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
		parsedDate, parseErr := time.Parse("2006-01-02", qSurveyDate)
		if parseErr != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD.")
			return
		}
		survey, err := s.surveyService.GetSurveyBySurveyDate(parsedDate)
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
	isadmin, _, errToken := s.userService.IsAdmin(token)
	if errToken != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errToken.Error())
		return
	}

	if !isadmin {
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

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, errToken := s.userService.IsAdmin(token)
	if errToken != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errToken.Error())
		return
	}
	if !isadmin {
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
	surveyID := c.Param("id")

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, errToken := s.userService.IsAdmin(token)
	if errToken != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errToken.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can delete survey!")
		return
	}

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
