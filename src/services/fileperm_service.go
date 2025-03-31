package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/repo"
)

type filePermServiceImpl struct {
	filePermRepo repo.FilePermRepo
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
func (f *filePermServiceImpl) RejectFileAccess(request request.FilePermRequest) error {
	err := f.filePermRepo.RejectAccess(request.UserID, request.FileID)
	if err != nil {
		return err
	}
	return nil
}

// ApproveFileAccess implements FileUploadPermService.
func (f *filePermServiceImpl) ApproveFileAccess(request request.FilePermRequest) error {
	err := f.filePermRepo.ApproveAccess(request.UserID, request.FileID)
	if err != nil {
		return err
	}
	return nil
}

// RequestAccess implements FileUploadPermService.
func (f *filePermServiceImpl) RequestAccess(request request.FilePermRequest) error {
	err := f.filePermRepo.RequestAccess(request.UserID, request.FileID)
	if err != nil {
		return err
	}
	return nil
}

func NewFilePermServceImpl(filePermRepo repo.FilePermRepo) FileUploadPermService {
	return &filePermServiceImpl{filePermRepo: filePermRepo}
}
