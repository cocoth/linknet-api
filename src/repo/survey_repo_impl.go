package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type SurveyRepoImpl struct {
	db *gorm.DB
}

// GetAllSurvey implements SurveyRepo.
func (s *SurveyRepoImpl) GetAllSurvey() ([]models.Survey, error) {
	var surveys []models.Survey
	if err := s.db.Find(&surveys).Error; err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveyByAddress implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByAddress(address string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "address = ?", address).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByCustomerName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByCustomerName(customerName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "customer_name = ?", customerName).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByFAT implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByFAT(fat string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "fat = ?", fat).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByFormNumber implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByFormNumber(formNumber string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "formn_number = ?", formNumber).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByID(id string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "id = ?", id).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByNodeFDT implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByNodeFDT(nodeFDT string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "node_fdt = ?", nodeFDT).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByQuestorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByQuestorName(questorName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "questor_name = ?", questorName).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyDate implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyDate(surveyDate string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "survey_date = ?", surveyDate).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyorID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyorID(surveyorID string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "surveyor_id = ?", surveyorID).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyorName(surveyorName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Joins("JOIN users ON users.id = surveys.surveyor_id").Where("users.name = ?", surveyorName).First(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByTitle implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByTitle(title string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "title = ?", title).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByImageID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByImageID(imageID string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.First(&survey, "image_id = ?", imageID).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveysByAddress implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByAddress(address string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Where("address LIKE ?", "%"+address+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByCustomerName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByCustomerName(customerName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Where("customer_name LIKE ?", "%"+customerName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByQuestorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByQuestorName(questorName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Where("questor_name LIKE ?", "%"+questorName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysBySurveyorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysBySurveyorName(surveyorName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Where("surveyor_name LIKE ?", "%"+surveyorName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByTitle implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByTitle(title string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Where("title LIKE ?", "%"+title+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// CreateSurvey implements SurveyRepo.
func (s *SurveyRepoImpl) CreateSurvey(survey models.Survey) (models.Survey, error) {
	if err := s.db.Create(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// UpdateSurvey implements SurveyRepo.
func (s *SurveyRepoImpl) UpdateSurvey(survey models.Survey) (models.Survey, error) {
	err := s.db.Model(&survey).Where("id = ?", survey.ID).Updates(survey).Error
	if err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// DeleteSurvey implements SurveyRepo.
func (s *SurveyRepoImpl) DeleteSurvey(id string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Where("id = ?", id).Delete(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}
func NewSurveyRepoImpl(db *gorm.DB) SurveyRepo {
	return &SurveyRepoImpl{db: db}
}
