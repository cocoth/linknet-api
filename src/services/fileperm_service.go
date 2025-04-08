package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
)

type filePermServiceImpl struct {
	filePermRepo repo.FilePermRepo
}

func sendFilePermService(data models.FileAccessRequest, err error) (response.FilePermResponse, error) {
	if err != nil {
		return response.FilePermResponse{}, err
	}
	return response.FilePermResponse{
		ID:        data.ID,
		FileID:    data.FileID,
		UserID:    data.UserID,
		Approved:  data.Approved,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}, nil
}

// CheckAccess implements FileUploadPermService.
func (f *filePermServiceImpl) CheckAccess(request request.FilePermRequest) (bool, error) {
	access, err := f.filePermRepo.CheckAccess(request.UserID, request.FileID)
	if err != nil {
		return false, err
	}
	return access, nil
}

// RejectFileAccess implements FileUploadPermService.
func (f *filePermServiceImpl) RejectFileAccess(request request.FilePermRequest) (response.FilePermResponse, error) {
	perm, err := f.filePermRepo.RejectAccess(request.UserID, request.FileID)
	return sendFilePermService(perm, err)
}

// ApproveFileAccess implements FileUploadPermService.
func (f *filePermServiceImpl) ApproveFileAccess(request request.FilePermRequest) (response.FilePermResponse, error) {
	perm, err := f.filePermRepo.ApproveAccess(request.UserID, request.FileID)
	return sendFilePermService(perm, err)

}

// RequestAccess implements FileUploadPermService.
func (f *filePermServiceImpl) RequestAccess(request request.FilePermRequest) (response.FilePermResponse, error) {
	perm, err := f.filePermRepo.RequestAccess(request.UserID, request.FileID)
	return sendFilePermService(perm, err)

}

func NewFilePermServceImpl(filePermRepo repo.FilePermRepo) FileUploadPermService {
	return &filePermServiceImpl{filePermRepo: filePermRepo}
}
