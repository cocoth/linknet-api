package services

import (
	"time"

	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
)

type SurveyServiceImpl struct {
	surveyRepo repo.SurveyRepo
}

func sendSurveyResponse(survey models.Survey, err error) (response.SurveyResponse, error) {
	if err != nil {
		return response.SurveyResponse{}, err
	}

	var surveyors []response.SurveyorLinkResponse
	for _, surveyor := range survey.Surveyors {
		surveyors = append(surveyors, response.SurveyorLinkResponse{
			ID:         surveyor.ID,
			SurveyID:   surveyor.SurveyID,
			SurveyorID: surveyor.SurveyorID,
			CreatedAt:  surveyor.CreatedAt,
			UpdatedAt:  surveyor.UpdatedAt,
			DeletedAt:  surveyor.DeletedAt,
		})
	}
	return response.SurveyResponse{
		ID:           survey.ID,
		Title:        survey.Title,
		FormNumber:   survey.FormNumber,
		QuestorName:  survey.QuestorName,
		FAT:          survey.FAT,
		CustomerName: survey.CustomerName,
		Address:      survey.Address,
		NodeFDT:      survey.NodeFDT,
		SurveyDate:   survey.SurveyDate,
		Surveyors:    surveyors,
		ImageID:      survey.ImageID,
		CreatedAt:    survey.CreatedAt,
		UpdatedAt:    survey.UpdatedAt,
		DeletedAt:    survey.DeletedAt,
	}, nil
}

func sendSurveysResponse(surveyModel []models.Survey, err error) ([]response.SurveyResponse, error) {
	if err != nil {
		return nil, err
	}
	surveys := make([]response.SurveyResponse, 0, len(surveyModel))
	for _, survey := range surveyModel {
		var surveyors []response.SurveyorLinkResponse
		for _, surveyor := range survey.Surveyors {
			surveyors = append(surveyors, response.SurveyorLinkResponse{
				ID:         surveyor.ID,
				SurveyID:   surveyor.SurveyID,
				SurveyorID: surveyor.SurveyorID,
				CreatedAt:  surveyor.CreatedAt,
				UpdatedAt:  surveyor.UpdatedAt,
				DeletedAt:  surveyor.DeletedAt,
			})
		}
		surveys = append(surveys, response.SurveyResponse{
			ID:           survey.ID,
			Title:        survey.Title,
			FormNumber:   survey.FormNumber,
			QuestorName:  survey.QuestorName,
			FAT:          survey.FAT,
			CustomerName: survey.CustomerName,
			Address:      survey.Address,
			NodeFDT:      survey.NodeFDT,
			SurveyDate:   survey.SurveyDate,
			Surveyors:    surveyors,
			ImageID:      survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}
	return surveys, nil

}

// ViewSurveyAndReportsByID implements SurveyService.
func (s *SurveyServiceImpl) ViewSurveyAndReportsByID(id string) (response.SurveyReportView, error) {
	survey, err := s.surveyRepo.ViewSurveyAndReportsByID(id)
	if err != nil {
		return response.SurveyReportView{}, err
	}

	var surveyors []string
	for _, surveyor := range survey.Surveyors {
		if surveyor.Surveyor.Name != "" {
			surveyors = append(surveyors, surveyor.Surveyor.Name)
		} else {
			surveyors = append(surveyors, surveyor.Surveyor.CallSign)
		}
	}

	return response.SurveyReportView{
		FormNumber:   survey.FormNumber,
		QuestorName:  survey.QuestorName,
		Fat:          survey.FAT,
		CustomerName: survey.CustomerName,
		Address:      survey.Address,
		NodeFDT:      survey.NodeFDT,
		SurveyDate:   survey.SurveyDate,
		Status:       survey.SurveyReport.Status,
		Remark:       survey.SurveyReport.Remark,
		Surveyors:    surveyors,
	}, nil
}

// GetSurveysWithFilters implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysWithFilters(filters map[string]interface{}) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysWithFilters(filters)
	return sendSurveysResponse(survey, err)
}

// GetAllSurvey implements SurveyService.
func (s *SurveyServiceImpl) GetAllSurvey() ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetAllSurvey()
	return sendSurveysResponse(survey, err)
}

// GetSurveyByAddress implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByAddress(address string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByAddress(address)
	return sendSurveyResponse(survey, err)

}

// GetSurveyByCustomerName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByCustomerName(customerName string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByCustomerName(customerName)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByFAT implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByFAT(fat string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByFAT(fat)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByFormNumber implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByFormNumber(formNumber string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByFormNumber(formNumber)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByID implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByID(id string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByID(id)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByNodeFDT implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByNodeFDT(nodeFDT string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByNodeFDT(nodeFDT)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByQuestorName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByQuestorName(questorName string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByQuestorName(questorName)
	return sendSurveyResponse(survey, err)
}

// GetSurveyBySurveyDate implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyBySurveyDate(surveyDate time.Time) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyBySurveyDate(surveyDate)
	return sendSurveyResponse(survey, err)
}

// GetSurveyBySurveyorID implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyBySurveyorID(surveyorID string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyBySurveyorID(surveyorID)
	return sendSurveyResponse(survey, err)
}

// GetSurveysBySurveyorName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysBySurveyorName(surveyorName string) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysBySurveyorName(surveyorName)
	return sendSurveysResponse(survey, err)
}

// GetSurveyByTitle implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByTitle(title string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByTitle(title)
	return sendSurveyResponse(survey, err)
}

// GetSurveyByImageID implements SurveyService.
func (s *SurveyServiceImpl) GetSurveyByImageID(imageID string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveyByImageID(imageID)
	return sendSurveyResponse(survey, err)
}

// GetSurveysByAddress implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByAddress(address string) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysByAddress(address)
	return sendSurveysResponse(survey, err)
}

// GetSurveysByCustomerName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByCustomerName(customerName string) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysByCustomerName(customerName)
	return sendSurveysResponse(survey, err)
}

// GetSurveysByQuestorName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByQuestorName(questorName string) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysByQuestorName(questorName)
	return sendSurveysResponse(survey, err)
}

// GetSurveysByTitle implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByTitle(title string) ([]response.SurveyResponse, error) {
	survey, err := s.surveyRepo.GetSurveysByTitle(title)
	return sendSurveysResponse(survey, err)
}

// CreateSurvey implements SurveyService.
func (s *SurveyServiceImpl) CreateSurvey(survey request.SurveyRequest) (response.SurveyResponse, error) {
	// var survey models.Survey

	survey.Title = utils.SanitizeString(survey.Title)
	survey.FormNumber = utils.SanitizeString(survey.FormNumber)
	survey.QuestorName = utils.SanitizeString(survey.QuestorName)
	survey.FAT = utils.SanitizeString(survey.FAT)
	survey.CustomerName = utils.SanitizeString(survey.CustomerName)
	survey.Address = utils.SanitizeString(survey.Address)
	survey.NodeFDT = utils.SanitizeString(survey.NodeFDT)

	var surveyors []models.SurveyorLink
	for _, surveyor := range survey.Surveyors {
		surveyorModel := models.SurveyorLink{
			SurveyID:   surveyor.SurveyID,
			SurveyorID: surveyor.SurveyorID,
		}
		surveyors = append(surveyors, surveyorModel)
	}

	surveyModel := models.Survey{
		Title:        survey.Title,
		FormNumber:   survey.FormNumber,
		QuestorName:  survey.QuestorName,
		FAT:          survey.FAT,
		CustomerName: survey.CustomerName,
		Address:      survey.Address,
		NodeFDT:      survey.NodeFDT,
		SurveyDate:   survey.SurveyDate,
		Surveyors:    surveyors,
		ImageID:      &survey.ImageID,
	}

	surveyCreate, err := s.surveyRepo.CreateSurvey(surveyModel)
	return sendSurveyResponse(surveyCreate, err)

}

// UpdateSurvey implements SurveyService.
func (s *SurveyServiceImpl) UpdateSurvey(id string, survey request.UpdateSurveyRequest) (response.SurveyResponse, error) {
	surveyData, err := s.surveyRepo.GetSurveyByID(id)
	if err != nil {
		return response.SurveyResponse{}, err
	}

	if survey.Title != nil {
		sanitizedTitle := utils.SanitizeString(*survey.Title)
		surveyData.Title = sanitizedTitle
	}
	if survey.Address != nil {
		sanitizedAddress := utils.SanitizeString(*survey.Address)
		surveyData.Address = sanitizedAddress
	}
	if survey.CustomerName != nil {
		sanitizedCustomerName := utils.SanitizeString(*survey.CustomerName)
		surveyData.CustomerName = sanitizedCustomerName
	}
	if survey.FAT != nil {
		sanitizedFAT := utils.SanitizeString(*survey.FAT)
		surveyData.FAT = sanitizedFAT
	}
	if survey.FormNumber != nil {
		sanitizedFormNumber := utils.SanitizeString(*survey.FormNumber)
		surveyData.FormNumber = sanitizedFormNumber
	}
	if survey.NodeFDT != nil {
		sanitizedNodeFDT := utils.SanitizeString(*survey.NodeFDT)
		surveyData.NodeFDT = sanitizedNodeFDT
	}
	if survey.QuestorName != nil {
		sanitizedQuestorName := utils.SanitizeString(*survey.QuestorName)
		surveyData.QuestorName = sanitizedQuestorName
	}
	if survey.SurveyDate != nil {
		surveyData.SurveyDate = *survey.SurveyDate
	}
	if survey.Surveyors != nil {
		var surveyors []models.SurveyorLink
		for _, surveyor := range survey.Surveyors {
			sanitizedSurveyorID := utils.SanitizeString(surveyor.SurveyorID)
			surveyors = append(surveyors, models.SurveyorLink{
				SurveyID:   id,
				SurveyorID: sanitizedSurveyorID,
			})
		}
		surveyData.Surveyors = surveyors
	}
	if survey.ImageID != nil {
		surveyData.ImageID = survey.ImageID
	}

	updatedSurvey, err := s.surveyRepo.UpdateSurvey(surveyData)
	return sendSurveyResponse(updatedSurvey, err)
}

// DeleteSurvey implements SurveyService.
func (s *SurveyServiceImpl) DeleteSurvey(id string) (response.SurveyResponse, error) {
	survey, err := s.surveyRepo.DeleteSurvey(id)
	return sendSurveyResponse(survey, err)
}

func NewSurveyServiceImpl(surveyRepo repo.SurveyRepo) SurveyService {
	return &SurveyServiceImpl{
		surveyRepo: surveyRepo,
	}
}
