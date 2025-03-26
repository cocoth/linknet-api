package services

import (
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
		SurveyorID:   survey.SurveyorID,
		ImageID:      *survey.ImageID,
		CreatedAt:    survey.CreatedAt,
		UpdatedAt:    survey.UpdatedAt,
		DeletedAt:    survey.DeletedAt,
	}, nil
}

// GetAllSurvey implements SurveyService.
func (s *SurveyServiceImpl) GetAllSurvey() ([]response.SurveyResponse, error) {
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetAllSurvey()
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
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
func (s *SurveyServiceImpl) GetSurveyBySurveyDate(surveyDate string) (response.SurveyResponse, error) {
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
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetSurveysBySurveyorName(surveyorName)
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
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
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetSurveysByAddress(address)
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
}

// GetSurveysByCustomerName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByCustomerName(customerName string) ([]response.SurveyResponse, error) {
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetSurveysByCustomerName(customerName)
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
}

// GetSurveysByQuestorName implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByQuestorName(questorName string) ([]response.SurveyResponse, error) {
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetSurveysByQuestorName(questorName)
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
}

// GetSurveysByTitle implements SurveyService.
func (s *SurveyServiceImpl) GetSurveysByTitle(title string) ([]response.SurveyResponse, error) {
	var surveys []response.SurveyResponse

	survey, err := s.surveyRepo.GetSurveysByTitle(title)
	if err != nil {
		return nil, err
	}

	for _, survey := range survey {
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
			SurveyorID:   survey.SurveyorID,
			ImageID:      *survey.ImageID,
			CreatedAt:    survey.CreatedAt,
			UpdatedAt:    survey.UpdatedAt,
			DeletedAt:    survey.DeletedAt,
		})
	}

	return surveys, nil
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
	survey.SurveyorID = utils.SanitizeString(survey.SurveyorID)

	surveyModel := models.Survey{
		Title:        survey.Title,
		FormNumber:   survey.FormNumber,
		QuestorName:  survey.QuestorName,
		FAT:          survey.FAT,
		CustomerName: survey.CustomerName,
		Address:      survey.Address,
		NodeFDT:      survey.NodeFDT,
		SurveyDate:   survey.SurveyDate,
		SurveyorID:   survey.SurveyorID,
		ImageID:      &survey.ImageID,
	}

	surveyCreate, err := s.surveyRepo.CreateSurvey(surveyModel)
	return sendSurveyResponse(surveyCreate, err)

}

// UpdateSurvey implements SurveyService.
func (s *SurveyServiceImpl) UpdateSurvey(id string, survey request.UpdateSurveyRequest) (response.SurveyResponse, error) {
	panic("unimplemented")
}

// DeleteSurvey implements SurveyService.
func (s *SurveyServiceImpl) DeleteSurvey(id string) (response.SurveyResponse, error) {
	panic("unimplemented")
}

func NewSurveyServiceImpl(surveyRepo repo.SurveyRepo) SurveyService {
	return &SurveyServiceImpl{
		surveyRepo: surveyRepo,
	}
}
