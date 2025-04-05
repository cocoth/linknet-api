package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
)

type surveyReportServiceImpl struct {
	surveyReportRepo repo.SurveyReportRepo
}

func sendSurveyReport(data models.SurveyReport, err error) (response.SurveyReportResponse, error) {
	if err != nil {
		return response.SurveyReportResponse{}, err
	}
	var imageID string
	if data.ImageID != nil {
		imageID = *data.ImageID
	}
	return response.SurveyReportResponse{
		ID:        data.ID,
		SurveyID:  data.SurveyID,
		Status:    data.Status,
		Remark:    data.Remark,
		ImageID:   &imageID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}, nil
}

func sendSurveysReport(data []models.SurveyReport, err error) ([]response.SurveyReportResponse, error) {
	if err != nil {
		return nil, err
	}

	reports := make([]response.SurveyReportResponse, 0, len(data))
	for _, report := range data {
		var imageID string

		if report.ImageID != nil {
			imageID = *report.ImageID
		}
		reports = append(reports, response.SurveyReportResponse{
			ID:        report.ID,
			SurveyID:  report.SurveyID,
			Status:    report.Status,
			Remark:    report.Remark,
			ImageID:   &imageID,
			CreatedAt: report.CreatedAt,
			UpdatedAt: report.UpdatedAt,
			DeletedAt: report.DeletedAt,
		})
	}

	return reports, nil
}

// GetSurveysUploadWithFilters implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportsUploadWithFilters(filters map[string]interface{}) ([]response.SurveyReportResponse, error) {
	reports, err := s.surveyReportRepo.GetSurveyReportsUploadWithFilters(filters)
	return sendSurveysReport(reports, err)
}

// CreateSurveyReport implements SurveyReportService.
func (s *surveyReportServiceImpl) CreateSurveyReport(surveyReport request.SurveyReportRequest) (response.SurveyReportResponse, error) {
	surveyReport.Status = utils.SanitizeString(surveyReport.Status)
	surveyReport.Remark = utils.SanitizeString(surveyReport.Remark)

	report := models.SurveyReport{
		SurveyID: surveyReport.SurveyID,
		Status:   surveyReport.Status,
		Remark:   surveyReport.Remark,
		ImageID:  surveyReport.ImageID,
	}

	reportRes, err := s.surveyReportRepo.CreateSurveyReport(report)
	return sendSurveyReport(reportRes, err)

}

// DeleteSurveyReport implements SurveyReportService.
func (s *surveyReportServiceImpl) DeleteSurveyReport(id string) (response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.DeleteSurveyReport(id)
	return sendSurveyReport(report, err)
}

// GetAllSurveyReport implements SurveyReportService.
func (s *surveyReportServiceImpl) GetAllSurveyReport() ([]response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetAllSurveyReport()
	return sendSurveysReport(report, err)
}

// GetSurveyBySurveyID implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyBySurveyID(surveyID string) (response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveyReportBySurveyID(surveyID)
	return sendSurveyReport(report, err)
}

// GetSurveyReportByID implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportByID(id string) (response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveyReportByID(id)
	return sendSurveyReport(report, err)
}

// GetSurveyReportByRemark implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportByRemark(remark string) (response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveyReportByRemark(remark)
	return sendSurveyReport(report, err)
}

// GetSurveyReportByStatus implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportByStatus(status string) (response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveyReportByStatus(status)
	return sendSurveyReport(report, err)
}

// GetSurveysReportByRemark implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportsByRemark(remark string) ([]response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveysReportByRemark(remark)
	return sendSurveysReport(report, err)
}

// GetSurveysReportByStatus implements SurveyReportService.
func (s *surveyReportServiceImpl) GetSurveyReportsByStatus(status string) ([]response.SurveyReportResponse, error) {
	report, err := s.surveyReportRepo.GetSurveysReportByStatus(status)
	return sendSurveysReport(report, err)
}

// UpdateSurveyReport implements SurveyReportService.
func (s *surveyReportServiceImpl) UpdateSurveyReport(id string, surveyReport request.UpdateSurveyReportRequest) (response.SurveyReportResponse, error) {
	dataReport, err := s.surveyReportRepo.GetSurveyReportByID(id)
	if err != nil {
		return response.SurveyReportResponse{}, err
	}

	if surveyReport.Status != nil {
		sanitizedStatus := utils.SanitizeString(*surveyReport.Status)
		dataReport.Status = sanitizedStatus
	}

	if surveyReport.Remark != nil {
		sanitizedRemark := utils.SanitizeString(*surveyReport.Remark)
		dataReport.Remark = sanitizedRemark
	}

	if surveyReport.ImageID != nil {
		dataReport.ImageID = surveyReport.ImageID
	}

	updatedReport, err := s.surveyReportRepo.UpdateSurveyReport(dataReport)
	return sendSurveyReport(updatedReport, err)

}

func NewSurveyReportServiceImpl(report repo.SurveyReportRepo) SurveyReportService {
	return &surveyReportServiceImpl{
		surveyReportRepo: report,
	}
}
