package repo

import (
	"time"

	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type SurveyRepoImpl struct {
	db *gorm.DB
}

// GetUsersWithFilters implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysWithFilters(filters map[string]interface{}) ([]models.Survey, error) {
	var surveys []models.Survey
	query := s.db.Preload("Surveyors")

	if id, ok := filters["id"]; ok {
		query = query.Where("id = ?", id)
	}
	if title, ok := filters["title"]; ok {
		query = query.Where("title LIKE ?", "%"+title.(string)+"%")
	}
	if form_number, ok := filters["form_number"]; ok {
		query = query.Where("call_sign = ?", form_number)
	}
	if questor_name, ok := filters["questor_name"]; ok {
		query = query.Where("questor_name = ?", questor_name)
	}
	if fat, ok := filters["fat"]; ok {
		query = query.Where("fat LIKE ?", "%"+fat.(string)+"%")
	}
	if customer_name, ok := filters["customer_name"]; ok {
		query = query.Where("customer_name LIKE ?", "%"+customer_name.(string)+"%")
	}
	if address, ok := filters["address"]; ok {
		query = query.Where("address LIKE ?", "%"+address.(string)+"%")
	}
	if node_fdt, ok := filters["node_fdt"]; ok {
		query = query.Where("node_fdt = ?", node_fdt)
	}
	if survey_date, ok := filters["survey_date"]; ok {
		query = query.Where("survey_date = ?", survey_date.(string))
	}
	if surveyor_id, ok := filters["surveyor_id"]; ok {
		query = query.Joins("JOIN surveyor_links ON surveyor_links.survey_id = surveys.id").
			Where("surveyor_links.surveyor_id = ?", surveyor_id)
	}
	if image_id, ok := filters["image_id"]; ok {
		query = query.Where("image_id = ?", image_id.(string))
	}

	// for key, value := range filters {
	// 	query = query.Where(key+" = ?", value)
	// }

	if err := query.Find(&surveys).Error; err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetAllSurveyWithPreload implements SurveyRepo.
func (s *SurveyRepoImpl) GetAllSurveyWithPreload(preload string) ([]models.Survey, error) {
	var surveys []models.Survey
	if err := s.db.Preload(preload).Find(&surveys).Error; err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetAllSurvey implements SurveyRepo.
func (s *SurveyRepoImpl) GetAllSurvey() ([]models.Survey, error) {
	surveys, err := s.GetAllSurveyWithPreload("Surveyors")
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveyByAddress implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByAddress(address string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "address = ?", address).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByCustomerName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByCustomerName(customerName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "customer_name = ?", customerName).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByFAT implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByFAT(fat string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "fat = ?", fat).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByFormNumber implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByFormNumber(formNumber string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "form_number = ?", formNumber).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByID(id string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "id = ?", id).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByNodeFDT implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByNodeFDT(nodeFDT string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "node_fdt = ?", nodeFDT).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByQuestorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByQuestorName(questorName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "questor_name = ?", questorName).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyDate implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyDate(surveyDate time.Time) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "survey_date = ?", surveyDate).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyorID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyorID(surveyorID string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").Joins("JOIN surveyor_links ON surveyor_links.survey_id = surveys.id").
		Where("surveyor_links.surveyor_id = ?", surveyorID).First(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyBySurveyorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyBySurveyorName(surveyorName string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").Joins("JOIN surveyor_links ON surveyor_links.survey_id = surveys.id").
		Joins("JOIN users ON users.id = surveyor_links.surveyor_id").
		Where("users.name = ?", surveyorName).First(&survey).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByTitle implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByTitle(title string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "title = ?", title).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveyByImageID implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveyByImageID(imageID string) (models.Survey, error) {
	var survey models.Survey
	if err := s.db.Preload("Surveyors").First(&survey, "image_id = ?", imageID).Error; err != nil {
		return models.Survey{}, err
	}
	return survey, nil
}

// GetSurveysByAddress implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByAddress(address string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Preload("Surveyors").Where("address LIKE ?", "%"+address+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByCustomerName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByCustomerName(customerName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Preload("Surveyors").Where("customer_name LIKE ?", "%"+customerName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByQuestorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByQuestorName(questorName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Preload("Surveyors").Where("questor_name LIKE ?", "%"+questorName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysBySurveyorName implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysBySurveyorName(surveyorName string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Preload("Surveyors").Joins("JOIN surveyor_links ON surveyor_links.survey_id = surveys.id").
		Joins("JOIN users ON users.id = surveyor_links.surveyor_id").
		Where("users.name LIKE ?", "%"+surveyorName+"%").Find(&surveys).Error
	if err != nil {
		return nil, err
	}
	return surveys, nil
}

// GetSurveysByTitle implements SurveyRepo.
func (s *SurveyRepoImpl) GetSurveysByTitle(title string) ([]models.Survey, error) {
	var surveys []models.Survey

	err := s.db.Preload("Surveyors").Where("title LIKE ?", "%"+title+"%").Find(&surveys).Error
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
