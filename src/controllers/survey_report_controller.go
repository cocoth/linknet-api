package controllers

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type SurveyReportController struct {
	reportService services.SurveyReportService
	userService   services.UserService
}

func NewSurveyReportController(reportService services.SurveyReportService, userService services.UserService) *SurveyReportController {
	return &SurveyReportController{
		reportService: reportService,
		userService:   userService,
	}
}

func (r *SurveyReportController) GetAllReport(c *gin.Context) {

	qID := c.Query("id")
	qStatus := c.Query("status")
	qRemark := c.Query("remark")

	filters := map[string]interface{}{}

	if qID != "" {
		filters["id"] = qID
	}
	if qStatus != "" {
		filters["status"] = qStatus
	}
	if qRemark != "" {
		filters["remark"] = qRemark
	}

	if qID == "" && qStatus == "" && qRemark == "" {
		helper.RespondWithError(c, http.StatusNotFound, "No query parameters provided")
		return
	}

	reports, err := r.reportService.GetSurveyReportsUploadWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(reports) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No Reports found with the given filters")
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, reports)
}

func (r *SurveyReportController) CreateSurveyReport(c *gin.Context) {
	var surveyRequest request.SurveyReportRequest
	if err := c.ShouldBindJSON(&surveyRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	surveyReport, err := r.reportService.CreateSurveyReport(surveyRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusCreated, surveyReport)
}

func (r *SurveyReportController) UpdateSurveyReport(c *gin.Context) {
	var surveyRequest request.UpdateSurveyReportRequest
	if err := c.ShouldBindJSON(&surveyRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	surveyID := c.Param("id")
	surveyReport, err := r.reportService.UpdateSurveyReport(surveyID, surveyRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, surveyReport)
}

func (r *SurveyReportController) DeleteSurveyReport(c *gin.Context) {
	surveyID := c.Param("id")
	surveyReport, err := r.reportService.DeleteSurveyReport(surveyID)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, surveyReport)
}
