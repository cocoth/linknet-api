package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
)

type fileUploadServiceImpl struct {
	fileRepo repo.FileUploadRepo
}

func sendFileUploadResponse(data models.FileUpload, err error) (response.FileUploadResponse, error) {
	if err != nil {
		return response.FileUploadResponse{}, err
	}
	return response.FileUploadResponse{
		ID:        data.ID,
		FileName:  data.FileName,
		FileType:  data.FileType,
		FileUri:   data.FileUri,
		FileHash:  data.FileHash,
		AuthorID:  data.AuthorID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}, nil
}

// UploadFile implements FileUploadService.
func (f *fileUploadServiceImpl) UploadFile(file request.FileUploadRequest) (response.FileUploadResponse, error) {
	// Save the file using the repository
	fileRes, err := f.fileRepo.UploadFile(models.FileUpload{
		FileName: file.FileName,
		FileType: file.FileType,
		FileUri:  file.FileUri,
		FileHash: file.FileHash,
		AuthorID: &file.AuthorID,
	})
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(fileRes, nil)
}

// GetAllFileUpload implements FileUploadService.
func (f *fileUploadServiceImpl) GetAllFileUpload() ([]response.FileUploadResponse, error) {
	files, err := f.fileRepo.GetAllFileUpload()
	if err != nil {
		return nil, err
	}

	var fileResponses []response.FileUploadResponse
	for _, file := range files {
		fileResponses = append(fileResponses, response.FileUploadResponse{
			ID:        file.ID,
			FileName:  file.FileName,
			FileType:  file.FileType,
			FileUri:   file.FileUri,
			FileHash:  file.FileHash,
			AuthorID:  file.AuthorID,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
			DeletedAt: file.DeletedAt,
		})
	}

	return fileResponses, nil
}

// GetFileUploadByFileHash implements FileUploadService.
func (f *fileUploadServiceImpl) GetFileUploadByFileHash(fileHash string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.GetFileUploadByFileHash(fileHash)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

// GetFileUploadByFileID implements FileUploadService.
func (f *fileUploadServiceImpl) GetFileUploadByFileID(id string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.GetFileUploadByFileID(id)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

// GetFileUploadByFileName implements FileUploadService.
func (f *fileUploadServiceImpl) GetFileUploadByFileName(fileName string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.GetFileUploadByFileName(fileName)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

// GetFilesUploadByAuthorID implements FileUploadService.
func (f *fileUploadServiceImpl) GetFilesUploadByAuthorID(authorID string) ([]response.FileUploadResponse, error) {
	files, err := f.fileRepo.GetFilesUploadByAuthorID(authorID)
	if err != nil {
		return nil, err
	}

	var fileResponses []response.FileUploadResponse
	for _, file := range files {
		fileResponses = append(fileResponses, response.FileUploadResponse{
			ID:        file.ID,
			FileName:  file.FileName,
			FileType:  file.FileType,
			FileUri:   file.FileUri,
			FileHash:  file.FileHash,
			AuthorID:  file.AuthorID,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
			DeletedAt: file.DeletedAt,
		})
	}

	return fileResponses, nil
}

// GetFilesUploadByFileName implements FileUploadService.
func (f *fileUploadServiceImpl) GetFilesUploadByFileName(fileName string) ([]response.FileUploadResponse, error) {
	files, err := f.fileRepo.GetFilesUploadByFileName(fileName)
	if err != nil {
		return nil, err
	}

	var fileResponses []response.FileUploadResponse
	for _, file := range files {
		fileResponses = append(fileResponses, response.FileUploadResponse{
			ID:        file.ID,
			FileName:  file.FileName,
			FileType:  file.FileType,
			FileUri:   file.FileUri,
			FileHash:  file.FileHash,
			AuthorID:  file.AuthorID,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
			DeletedAt: file.DeletedAt,
		})
	}

	return fileResponses, nil
}

// UpdateFileUploadByAuthorID implements FileUploadService.
func (f *fileUploadServiceImpl) UpdateFileUpload(id string, fileUpdate request.FileUploadRequest) (response.FileUploadResponse, error) {
	var file models.FileUpload
	var err error
	file, err = f.fileRepo.GetFileUploadByFileID(id)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	if fileUpdate.FileName != "" {
		file.FileName = fileUpdate.FileName
	}
	if fileUpdate.FileUri != "" {
		file.FileUri = fileUpdate.FileUri
	}
	if fileUpdate.FileHash != "" {
		file.FileHash = fileUpdate.FileHash
	}
	if fileUpdate.AuthorID != "" {
		file.AuthorID = &fileUpdate.AuthorID
	}

	updatedFile, err := f.fileRepo.UpdateFileUpload(file)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(updatedFile, nil)
}

// DeleteFileUploadByFileHash implements FileUploadService.
func (f *fileUploadServiceImpl) DeleteFileUploadByFileHash(fileHash string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.DeleteFileUploadByFileHash(fileHash)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

// DeleteFileUploadByFileID implements FileUploadService.
func (f *fileUploadServiceImpl) DeleteFileUploadByFileID(id string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.DeleteFileUploadByFileID(id)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

// DeleteFileUploadByFileName implements FileUploadService.
func (f *fileUploadServiceImpl) DeleteFileUploadByFileName(fileName string) (response.FileUploadResponse, error) {
	file, err := f.fileRepo.DeleteFileUploadByFileName(fileName)
	if err != nil {
		return response.FileUploadResponse{}, err
	}

	return sendFileUploadResponse(file, nil)
}

func NewFileUploadServiceImpl(file repo.FileUploadRepo) FileUploadService {
	return &fileUploadServiceImpl{
		fileRepo: file,
	}

}
