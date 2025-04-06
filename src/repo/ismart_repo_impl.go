package repo

import (
	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type iSmartRepoImpl struct {
	db *gorm.DB
}

// CreateISmart implements ISmartRepo.
func (i *iSmartRepoImpl) CreateISmart(iSmart models.ISmart) (models.ISmart, error) {
	if err := i.db.Create(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// DeleteISmart implements ISmartRepo.
func (i *iSmartRepoImpl) DeleteISmart(id string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("id = ?", id).Delete(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetAllISmart implements ISmartRepo.
func (i *iSmartRepoImpl) GetAllISmart() ([]models.ISmart, error) {
	var iSmarts []models.ISmart
	if err := i.db.Find(&iSmarts).Error; err != nil {
		return nil, err
	}
	return iSmarts, nil
}

// GetISmartByAddress implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartByAddress(address string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("address = ?", address).First(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetISmartByCoordinate implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartByCoordinate(coordinate string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("coordinate = ?", coordinate).First(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetISmartByFiberNode implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartByFiberNode(fiberNode string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("fiber_node = ?", fiberNode).First(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetISmartByID implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartByID(id string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("id = ?", id).First(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetISmartByStreet implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartByStreet(street string) (models.ISmart, error) {
	var iSmart models.ISmart
	if err := i.db.Where("street = ?", street).First(&iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

// GetISmartsWithFilters implements ISmartRepo.
func (i *iSmartRepoImpl) GetISmartsWithFilters(filters map[string]interface{}) ([]models.ISmart, error) {
	var iSmarts []models.ISmart
	query := i.db.Model(&models.ISmart{})

	if id, ok := filters["id"]; ok {
		query = query.Where("id = ?", id.(string))
	}
	if fiber_node, ok := filters["fiber_node"]; ok {
		query = query.Where("fiber_node = ?", fiber_node.(string))
	}
	if address, ok := filters["address"]; ok {
		query = query.Where("address LIKE ?", "%"+address.(string)+"%")
	}
	if coordinate, ok := filters["coordinate"]; ok {
		query = query.Where("coordinate = ?", coordinate.(string))
	}
	if street, ok := filters["street"]; ok {
		query = query.Where("street LIKE ?", "%"+street.(string)+"%")
	}

	if err := query.Find(&iSmarts).Error; err != nil {
		return nil, err
	}
	return iSmarts, nil
}

// UpdateISmart implements ISmartRepo.
func (i *iSmartRepoImpl) UpdateISmart(iSmart models.ISmart) (models.ISmart, error) {
	if err := i.db.Model(&iSmart).Where("id = ?", iSmart.ID).Updates(iSmart).Error; err != nil {
		return models.ISmart{}, err
	}
	return iSmart, nil
}

func NewISmartRepoImpl(db *gorm.DB) ISmartRepo {
	return &iSmartRepoImpl{
		db: db,
	}
}
