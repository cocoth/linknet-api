package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type SurveyReportImpl struct {
	db *gorm.DB
}

// GetSurveysUploadWithFilters implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveyReportsUploadWithFilters(filters map[string]interface{}) ([]models.SurveyReport, error) {
	var reports []models.SurveyReport

	query := s.db.Model(&models.SurveyReport{})

	if id, ok := filters["id"]; ok {
		query = query.Where("id = ?", id)
	}
	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if remark, ok := filters["remark"]; ok {
		query = query.Where("remark LIKE ?", "%"+remark.(string)+"%")
	}

	err := query.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// GetSurveysReportByRemark implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveysReportByRemark(remark string) ([]models.SurveyReport, error) {
	var reports []models.SurveyReport

	err := s.db.Where("remark LIKE ?", "%"+remark+"%").Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// GetSurveysReportByStatus implements SurveyReportRepo.
func (s *SurveyReportImpl) GetSurveysReportByStatus(status string) ([]models.SurveyReport, error) {
	var reports []models.SurveyReport

	err := s.db.Where("status LIKE ?", "%"+status+"%").Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
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

func NewSurveyReportRepoImpl(db *gorm.DB) SurveyReportRepo {
	return &SurveyReportImpl{db: db}
}
