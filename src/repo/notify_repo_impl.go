package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type NotifyRepoImpl struct {
	db *gorm.DB
}

// GetAllNotify implements NotifyRepo.
func (n *NotifyRepoImpl) GetAllNotify() ([]models.Notify, error) {
	var notifies []models.Notify

	if err := n.db.Find(&notifies).Error; err != nil {
		return nil, err
	}
	return notifies, nil
}

// GetNotifyByFileID implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByFileID(fileID string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "file_id = ?", fileID).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByID implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByID(id string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "id = ?", id).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByNotifyMessage implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByNotifyMessage(notifyMessage string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.Where("notify_message LIKE ?", "%"+notifyMessage+"%").First(&notify).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByNotifyStatus implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByNotifyStatus(notifyStatus string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "notify_status = ?", notifyStatus).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByNotifyType implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByNotifyType(notifyType string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "notify_type = ?", notifyType).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByUserID implements NotifyRepo.
func (n *NotifyRepoImpl) GetNotifyByUserID(userID string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "user_id = ?", userID).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// CreateNotify implements NotifyRepo.
func (n *NotifyRepoImpl) CreateNotify(notify models.Notify) (models.Notify, error) {
	if err := n.db.Create(&notify).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// UpdateNotify implements NotifyRepo.
func (n *NotifyRepoImpl) UpdateNotify(notify models.Notify) (models.Notify, error) {
	err := n.db.Model(&notify).Where("id = ?", notify.ID).Updates(notify).Error
	if err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// DeleteNotify implements NotifyRepo.
func (n *NotifyRepoImpl) DeleteNotify(id string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "id = ?", id).Error; err != nil {
		return models.Notify{}, err
	}

	if err := n.db.Delete(&notify).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

func NewNotifyRepoImpl(db *gorm.DB) NotifyRepo {
	return &NotifyRepoImpl{db: db}
}
