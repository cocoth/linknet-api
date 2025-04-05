package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
)

type NotifyServiceImpl struct {
	notifyRepo repo.NotifyRepo
}

func sendNotifyResponse(notifyModel models.Notify, err error) (response.NotifyResponse, error) {
	if err != nil {
		return response.NotifyResponse{}, err
	}

	var fileID string
	if notifyModel.FileID != nil {
		fileID = *notifyModel.FileID
	}
	return response.NotifyResponse{
		ID:            notifyModel.ID,
		UserID:        notifyModel.UserID,
		FileID:        fileID,
		NotifyType:    notifyModel.NotifyType,
		NotifyStatus:  notifyModel.NotifyStatus,
		NotifyMessage: notifyModel.NotifyMessage,
		IsRead:        notifyModel.IsRead,
		CreatedAt:     notifyModel.CreatedAt,
		UpdatedAt:     notifyModel.UpdatedAt,
		DeletedAt:     notifyModel.DeletedAt,
	}, nil
}

func sendNotifiesResponse(notifyModel []models.Notify, err error) ([]response.NotifyResponse, error) {
	if err != nil {
		return nil, err
	}

	notifyResponses := make([]response.NotifyResponse, 0, len(notifyModel))
	for _, notify := range notifyModel {
		notifyResponses = append(notifyResponses, response.NotifyResponse{
			ID:            notify.ID,
			UserID:        notify.UserID,
			FileID:        *notify.FileID,
			NotifyType:    notify.NotifyType,
			NotifyStatus:  notify.NotifyStatus,
			NotifyMessage: notify.NotifyMessage,
			IsRead:        notify.IsRead,
			CreatedAt:     notify.CreatedAt,
			UpdatedAt:     notify.UpdatedAt,
			DeletedAt:     notify.DeletedAt,
		})
	}

	return notifyResponses, nil
}

// GetNotifyWithFilters implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyWithFilters(filters map[string]interface{}) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyWithFilters(filters)
	return sendNotifiesResponse(notify, err)
}

// CreateNotify implements NotifyService.
func (n *NotifyServiceImpl) CreateNotify(notify request.NotifyRequest) (response.NotifyResponse, error) {
	notifyModel := models.Notify{
		UserID:        notify.UserID,
		FileID:        &notify.FileID,
		NotifyType:    notify.NotifyType,
		NotifyStatus:  notify.NotifyStatus,
		NotifyMessage: notify.NotifyMessage,
		IsRead:        notify.IsRead,
	}

	createdNotify, err := n.notifyRepo.CreateNotify(notifyModel)

	return sendNotifyResponse(createdNotify, err)
}

// DeleteNotify implements NotifyService.
func (n *NotifyServiceImpl) DeleteNotify(id string) (response.NotifyResponse, error) {
	notify, err := n.notifyRepo.DeleteNotify(id)
	return sendNotifyResponse(notify, err)
}

// GetAllNotify implements NotifyService.
func (n *NotifyServiceImpl) GetAllNotify() ([]response.NotifyResponse, error) {
	notifies, err := n.notifyRepo.GetAllNotify()
	if err != nil {
		return nil, err
	}

	var notifyResponses []response.NotifyResponse
	for _, notify := range notifies {
		notifyResponses = append(notifyResponses, response.NotifyResponse{
			ID:            notify.ID,
			UserID:        notify.UserID,
			FileID:        *notify.FileID,
			NotifyType:    notify.NotifyType,
			NotifyStatus:  notify.NotifyStatus,
			NotifyMessage: notify.NotifyMessage,
			IsRead:        notify.IsRead,
			CreatedAt:     notify.CreatedAt,
			UpdatedAt:     notify.UpdatedAt,
			DeletedAt:     notify.DeletedAt,
		})
	}

	return notifyResponses, nil

}

// GetNotifyByFileID implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByFileID(fileID string) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByFileID(fileID)
	return sendNotifiesResponse(notify, err)
}

// GetNotifyByID implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByID(id string) (response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByID(id)
	return sendNotifyResponse(notify, err)
}

// GetNotifyByNotifyMessage implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByNotifyMessage(notifyMessage string) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByNotifyMessage(notifyMessage)
	return sendNotifiesResponse(notify, err)
}

// GetNotifyByNotifyStatus implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByNotifyStatus(notifyStatus string) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByNotifyStatus(notifyStatus)
	return sendNotifiesResponse(notify, err)
}

// GetNotifyByNotifyType implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByNotifyType(notifyType string) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByNotifyType(notifyType)
	return sendNotifiesResponse(notify, err)
}

// GetNotifyByUserID implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByUserID(userID string) ([]response.NotifyResponse, error) {
	notify, err := n.notifyRepo.GetNotifyByUserID(userID)
	return sendNotifiesResponse(notify, err)
}

// GetNotifyByIsRead implements NotifyService.
func (n *NotifyServiceImpl) GetNotifyByIsRead(isRead bool) ([]response.NotifyResponse, error) {
	notifies, err := n.notifyRepo.GetNotifyByIsRead(isRead)
	if err != nil {
		return nil, err
	}

	var notifyResponses []response.NotifyResponse
	for _, notify := range notifies {
		notifyResponses = append(notifyResponses, response.NotifyResponse{
			ID:            notify.ID,
			UserID:        notify.UserID,
			FileID:        *notify.FileID,
			NotifyType:    notify.NotifyType,
			NotifyStatus:  notify.NotifyStatus,
			NotifyMessage: notify.NotifyMessage,
			IsRead:        notify.IsRead,
			CreatedAt:     notify.CreatedAt,
			UpdatedAt:     notify.UpdatedAt,
			DeletedAt:     notify.DeletedAt,
		})
	}

	return notifyResponses, nil
}

// UpdateNotify implements NotifyService.
func (n *NotifyServiceImpl) UpdateNotify(notify request.NotifyRequest) (response.NotifyResponse, error) {
	notifyModel := models.Notify{
		UserID:        notify.UserID,
		FileID:        &notify.FileID,
		NotifyType:    notify.NotifyType,
		NotifyStatus:  notify.NotifyStatus,
		NotifyMessage: notify.NotifyMessage,
		IsRead:        notify.IsRead,
	}

	updatedNotify, err := n.notifyRepo.UpdateNotify(notifyModel)

	return sendNotifyResponse(updatedNotify, err)
}

func NewNotifyServiceImpl(notifyRepo repo.NotifyRepo) NotifyService {
	return &NotifyServiceImpl{
		notifyRepo: notifyRepo,
	}
}
