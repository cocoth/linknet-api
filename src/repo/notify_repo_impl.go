package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type notifyRepoImpl struct {
	db *gorm.DB
}

// GetNotifyWithFilters implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyWithFilters(filters map[string]interface{}) ([]models.Notify, error) {
	var notify []models.Notify
	query := n.db.Model(&models.Notify{})

	if id, ok := filters["id"]; ok {
		query = query.Where("id = ?", id.(string))
	}
	if user_id, ok := filters["user_id"]; ok {
		query = query.Where("user_id = ?", user_id.(string))
	}
	if file_id, ok := filters["file_id"]; ok {
		query = query.Where("file_id = ?", file_id.(string))
	}
	if notify_type, ok := filters["notify_type"]; ok {
		query = query.Where("notify_type = ?", notify_type.(string))
	}
	if notify_status, ok := filters["notify_status"]; ok {
		query = query.Where("notify_status = ?", notify_status.(string))
	}
	if notify_message, ok := filters["notify_message"]; ok {
		query = query.Where("notify_message LIKE ?", "%"+notify_message.(string)+"%")
	}
	if is_read, ok := filters["is_read"]; ok {
		query = query.Where("is_read = ?", is_read.(bool))
	}

	if err := query.Find(&notify).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetAllNotify implements NotifyRepo.
func (n *notifyRepoImpl) GetAllNotify() ([]models.Notify, error) {
	var notifies []models.Notify

	if err := n.db.Find(&notifies).Error; err != nil {
		return nil, err
	}
	return notifies, nil
}

// GetNotifyByFileID implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByFileID(fileID string) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Find(&notify, "file_id = ?", fileID).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetNotifyByID implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByID(id string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.First(&notify, "id = ?", id).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

// GetNotifyByNotifyMessage implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByNotifyMessage(notifyMessage string) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Where("notify_message LIKE ?", "%"+notifyMessage+"%").Find(&notify).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetNotifyByNotifyStatus implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByNotifyStatus(notifyStatus string) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Find(&notify, "notify_status = ?", notifyStatus).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetNotifyByNotifyType implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByNotifyType(notifyType string) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Find(&notify, "notify_type = ?", notifyType).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetNotifyByUserID implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByUserID(userID string) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Find(&notify, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// GetNotifyByIsRead implements NotifyRepo.
func (n *notifyRepoImpl) GetNotifyByIsRead(isRead bool) ([]models.Notify, error) {
	var notify []models.Notify

	if err := n.db.Find(&notify, "is_read = ?", isRead).Error; err != nil {
		return nil, err
	}
	return notify, nil
}

// CreateNotify implements NotifyRepo.
func (n *notifyRepoImpl) CreateNotify(notify models.Notify) (models.Notify, error) {
	err := n.db.Create(&notify).Error
	return notify, err
}

// UpdateNotify implements NotifyRepo.
func (n *notifyRepoImpl) UpdateNotify(notify models.Notify) (models.Notify, error) {
	err := n.db.Model(&notify).Where("id = ?", notify.ID).Updates(notify).Error
	return notify, err
}

// DeleteNotify implements NotifyRepo.
func (n *notifyRepoImpl) DeleteNotify(id string) (models.Notify, error) {
	var notify models.Notify

	if err := n.db.Where("id = ?", id).Delete(&notify).Error; err != nil {
		return models.Notify{}, err
	}
	return notify, nil
}

func NewNotifyRepoImpl(db *gorm.DB) NotifyRepo {
	return &notifyRepoImpl{db: db}
}
