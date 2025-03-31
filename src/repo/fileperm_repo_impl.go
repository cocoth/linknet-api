package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type filePermRepoImpl struct {
	db *gorm.DB
}

// RejctAccess implements FilePermRepo.
func (f *filePermRepoImpl) RejectAccess(userID string, fileID string) error {
	var perm models.FileAccessRequest

	return f.db.Model(&perm).Where("user_id = ? AND file_id = ?", userID, fileID).Update("approved", false).Error
}

// ApproveAccess implements FilePermRepo.
func (f *filePermRepoImpl) ApproveAccess(userID string, fileID string) error {
	var perm models.FileAccessRequest

	return f.db.Model(&perm).Where("user_id = ? AND file_id = ?", userID, fileID).Update("approved", true).Error
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
func (f *filePermRepoImpl) RequestAccess(userID string, fileID string) error {
	var perm models.FileAccessRequest

	// Check if the record exists
	err := f.db.Where("user_id = ? AND file_id = ?", userID, fileID).First(&perm).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new record if not found
			perm = models.FileAccessRequest{
				UserID:   userID,
				FileID:   fileID,
				Approved: true,
			}
			return f.db.Create(&perm).Error
		}
		// Return other errors
		return err
	}

	// Update the existing record
	perm.Approved = true
	return f.db.Save(&perm).Error
}

func NewFilePermRepoImpl(db *gorm.DB) FilePermRepo {
	return &filePermRepoImpl{db: db}
}
