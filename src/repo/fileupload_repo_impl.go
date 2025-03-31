package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type fileUploadRepoImpl struct {
	Db *gorm.DB
}

// UploadFile implements FileUploadRepo.
func (f *fileUploadRepoImpl) UploadFile(file models.FileUpload) (models.FileUpload, error) {
	if err := f.Db.Create(&file).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// GetAllFileUpload implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetAllFileUpload() ([]models.FileUpload, error) {
	var files []models.FileUpload
	if err := f.Db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// GetFileUploadByFileID implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFileUploadByFileID(id string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "id = ?", id).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// GetFileUploadByFileName implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFileUploadByFileName(fileName string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "file_name = ?", fileName).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// GetFileUploadByFileHash implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFileUploadByFileHash(fileHash string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "file_hash = ?", fileHash).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// GetFilesUploadByAuthorID implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFilesUploadByAuthorID(authorID string) ([]models.FileUpload, error) {
	var files []models.FileUpload
	if err := f.Db.Where("author_id = ?", authorID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// GetFilesUploadByAuthorName implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFilesUploadByAuthorName(authorName string) ([]models.FileUpload, error) {
	var files []models.FileUpload
	if err := f.Db.Where("author_name = ?", authorName).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// GetFilesUploadByFileName implements FileUploadRepo.
func (f *fileUploadRepoImpl) GetFilesUploadByFileName(fileName string) ([]models.FileUpload, error) {
	var files []models.FileUpload
	if err := f.Db.Where("file_name LIKE ?", "%"+fileName+"%").Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// UpdateFileUpload implements FileUploadRepo.
func (f *fileUploadRepoImpl) UpdateFileUpload(file models.FileUpload) (models.FileUpload, error) {
	updatedFile := map[string]interface{}{}

	if file.FileName != "" {
		updatedFile["file_name"] = file.FileName
	}
	if file.FileHash != "" {
		updatedFile["file_hash"] = file.FileHash
	}
	if file.FileUri != "" {
		updatedFile["file_uri"] = file.FileUri
	}
	if file.AuthorID != nil {
		updatedFile["author_id"] = file.AuthorID
	} else {
		return file, nil
	}

	if len(updatedFile) == 0 {
		return file, nil
	}

	result := f.Db.Model(&file).Where("id = ?", file.ID).Updates(updatedFile)
	if result.Error != nil {
		return file, result.Error
	}

	err := f.Db.First(&file, "id = ?", file.ID).Error
	if err != nil {
		return file, err
	}
	return file, nil
}

// DeleteFileUploadByFileID implements FileUploadRepo.
func (f *fileUploadRepoImpl) DeleteFileUploadByFileID(id string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "id = ?", id).Error; err != nil {
		return models.FileUpload{}, err
	}
	if err := f.Db.Delete(&file).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// DeleteFileUploadByFileName implements FileUploadRepo.
func (f *fileUploadRepoImpl) DeleteFileUploadByFileName(fileName string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "file_name = ?", fileName).Error; err != nil {
		return models.FileUpload{}, err
	}
	if err := f.Db.Delete(&file).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

// DeleteFileUploadByFileHash implements FileUploadRepo.
func (f *fileUploadRepoImpl) DeleteFileUploadByFileHash(fileHash string) (models.FileUpload, error) {
	var file models.FileUpload
	if err := f.Db.First(&file, "file_hash = ?", fileHash).Error; err != nil {
		return models.FileUpload{}, err
	}
	if err := f.Db.Delete(&file).Error; err != nil {
		return models.FileUpload{}, err
	}
	return file, nil
}

func NewFileUploadRepoImpl(db *gorm.DB) FileUploadRepo {
	return &fileUploadRepoImpl{Db: db}

}
