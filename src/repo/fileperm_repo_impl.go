package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type filePermRepoImpl struct {
	db *gorm.DB
}

// RejctAccess implements FilePermRepo.
func (f *filePermRepoImpl) RejectAccess(userID string, fileID string) (models.FileAccessRequest, error) {
	var perm models.FileAccessRequest
	err := f.db.Model(&perm).
		Where("user_id = ? AND file_id = ?", userID, fileID).
		Updates(map[string]interface{}{
			"approved": false,
		}).Error
	if err != nil {
		return models.FileAccessRequest{}, err
	}

	errUpdate := f.db.Model(&perm).Where("user_id = ? AND file_id = ?", userID, fileID).Update("approved", false).Error
	if errUpdate != nil {
		return models.FileAccessRequest{}, errUpdate
	}
	return perm, nil
}

// ApproveAccess implements FilePermRepo.
func (f *filePermRepoImpl) ApproveAccess(userID string, fileID string) (models.FileAccessRequest, error) {
	var perm models.FileAccessRequest

	// Update status approved
	err := f.db.Model(&perm).
		Where("user_id = ? AND file_id = ?", userID, fileID).
		Updates(map[string]interface{}{
			"approved": true,
		}).Error
	if err != nil {
		return models.FileAccessRequest{}, err
	}

	errUpdate := f.db.Model(&perm).Where("user_id = ? AND file_id = ?", userID, fileID).First(&perm).Error
	if errUpdate != nil {
		return models.FileAccessRequest{}, errUpdate
	}
	return perm, nil
}

// CheckAccess implements FilePermRepo.
func (f *filePermRepoImpl) CheckAccess(userID string, fileID string) (bool, error) {
	var perm models.FileAccessRequest

	err := f.db.Where("user_id = ? AND file_id = ? AND approved = true", userID, fileID).First(&perm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return perm.Approved, nil
}

// RequestAccess implements FilePermRepo.
func (f *filePermRepoImpl) RequestAccess(userID string, fileID string) (models.FileAccessRequest, error) {
	var perm models.FileAccessRequest

	// Check if the record exists
	err := f.db.Where("user_id = ? AND file_id = ?", userID, fileID).First(&perm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new record if not found
			perm = models.FileAccessRequest{
				UserID:   userID,
				FileID:   fileID,
				Approved: false,
			}
			err = f.db.Create(&perm).Error
			return perm, err
		}
		// Return other errors
		return models.FileAccessRequest{}, err
	}

	// Update the existing record
	perm.Approved = false
	err = f.db.Save(&perm).Error
	return perm, err
}

func NewFilePermRepoImpl(db *gorm.DB) FilePermRepo {
	return &filePermRepoImpl{db: db}
}
