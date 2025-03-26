package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type SurveyReportImpl struct {
	db *gorm.DB
}

// GetAllSurveyReport implements SurveyReportRepo.
func (s *SurveyReportImpl) GetAllSurveyReport() ([]models.SurveyReport, error) {
	var reports []models.SurveyReport

	if err := s.db.Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// GetSurveyReportByID implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveyReportByID(id string) (models.SurveyReport, error) {
	var report models.SurveyReport

	if err := s.db.First(&report, "id = ?", id).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return report, nil
}

// GetSurveyReportByRemark implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveyReportByRemark(remark string) (models.SurveyReport, error) {
	var report models.SurveyReport

	if err := s.db.Where("remark LIKE ?", "%"+remark+"%").First(&report).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return report, nil
}

// GetSurveyReportByStatus implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveyReportByStatus(status string) (models.SurveyReport, error) {
	var report models.SurveyReport

	if err := s.db.First(&report, "status = ?", status).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return report, nil
}

// GetSurveyReportBySurveyID implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveyReportBySurveyID(surveyID string) (models.SurveyReport, error) {
	var report models.SurveyReport

	if err := s.db.First(&report, "survey_id = ?", surveyID).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return report, nil
}

// CreateSurveyReport implements SurveyReportRepo.
func (s *SurveyReportImpl) CreateSurveyReport(surveyReport models.SurveyReport) (models.SurveyReport, error) {
	if err := s.db.Create(&surveyReport).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return surveyReport, nil
}

// UpdateSurveyReport implements SurveyReportRepo.
func (s *SurveyReportImpl) UpdateSurveyReport(surveyReport models.SurveyReport) (models.SurveyReport, error) {
	err := s.db.Model(&surveyReport).Where("id = ?", surveyReport.ID).Updates(surveyReport).Error
	if err != nil {
		return models.SurveyReport{}, err
	}
	return surveyReport, nil
}

// DeleteSurveyReport implements SurveyReportRepo.
func (s *SurveyReportImpl) DeleteSurveyReport(id string) (models.SurveyReport, error) {
	var report models.SurveyReport

	if err := s.db.First(&report, "id = ?", id).Error; err != nil {
		return models.SurveyReport{}, err
	}

	if err := s.db.Delete(&report).Error; err != nil {
		return models.SurveyReport{}, err
	}
	return report, nil
}

func NewSurveyReportImpl(db *gorm.DB) SurveyReportRepo {
	return &SurveyReportImpl{db: db}
}
