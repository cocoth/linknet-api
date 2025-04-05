package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
)

type iSmartServiceImpl struct {
	ismartRepository repo.ISmartRepo
}

func sendISmartResponse(iSmart models.ISmart, err error) (response.ISmartResponse, error) {
	if err != nil {
		return response.ISmartResponse{}, err
	}

	return response.ISmartResponse{
		ID:           iSmart.ID,
		FiberNode:    iSmart.FiberNode,
		Address:      iSmart.Address,
		CustomerName: iSmart.CustomerName,
		Coordinate:   iSmart.Coordinate,
		Street:       iSmart.Street,
		CreatedAt:    iSmart.CreatedAt,
		UpdatedAt:    iSmart.UpdatedAt,
		DeletedAt:    iSmart.DeletedAt,
	}, nil
}

func sendISmartsResponse(iSmart []models.ISmart, err error) ([]response.ISmartResponse, error) {
	if err != nil {
		return nil, err
	}
	ismarts := make([]response.ISmartResponse, 0, len(iSmart))
	for _, iSmart := range iSmart {
		ismarts = append(ismarts, response.ISmartResponse{
			ID:           iSmart.ID,
			FiberNode:    iSmart.FiberNode,
			Address:      iSmart.Address,
			CustomerName: iSmart.CustomerName,
			Coordinate:   iSmart.Coordinate,
			Street:       iSmart.Street,
			CreatedAt:    iSmart.CreatedAt,
			UpdatedAt:    iSmart.UpdatedAt,
			DeletedAt:    iSmart.DeletedAt,
		})
	}
	return ismarts, nil
}

// CreateISmart implements ISmartService.
func (i *iSmartServiceImpl) CreateISmart(iSmart request.ISmartRequest) (response.ISmartResponse, error) {
	if iSmart.FiberNode == "" {
		sanitize := utils.SanitizeString(iSmart.FiberNode)
		iSmart.FiberNode = sanitize
	}
	if iSmart.Address == "" {
		sanitize := utils.SanitizeString(iSmart.Address)
		iSmart.Address = sanitize
	}
	if iSmart.CustomerName == "" {
		sanitize := utils.SanitizeString(iSmart.CustomerName)
		iSmart.CustomerName = sanitize
	}
	if iSmart.Coordinate == "" {
		sanitize := utils.SanitizeString(iSmart.Coordinate)
		iSmart.Coordinate = sanitize
	}
	if iSmart.Street == "" {
		sanitize := utils.SanitizeString(iSmart.Street)
		iSmart.Street = sanitize
	}
	// Convert request to model
	iSmartModel := models.ISmart{
		FiberNode:    iSmart.FiberNode,
		Address:      iSmart.Address,
		CustomerName: iSmart.CustomerName,
		Coordinate:   iSmart.Coordinate,
		Street:       iSmart.Street,
	}

	// Call repository to create ISmart
	createdISmart, err := i.ismartRepository.CreateISmart(iSmartModel)
	return sendISmartResponse(createdISmart, err)
}

// DeleteISmart implements ISmartService.
func (i *iSmartServiceImpl) DeleteISmart(id string) (response.ISmartResponse, error) {
	deletedISmart, err := i.ismartRepository.DeleteISmart(id)
	return sendISmartResponse(deletedISmart, err)
}

// GetAllISmart implements ISmartService.
func (i *iSmartServiceImpl) GetAllISmart() ([]response.ISmartResponse, error) {
	iSmarts, err := i.ismartRepository.GetAllISmart()
	return sendISmartsResponse(iSmarts, err)
}

// GetISmartByAddress implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByAddress(address string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByAddress(address)
	return sendISmartResponse(iSmart, err)
}

// GetISmartByCoordinate implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByCoordinate(coordinate string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByCoordinate(coordinate)
	return sendISmartResponse(iSmart, err)
}

// GetISmartByCustomerName implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByCustomerName(customerName string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByCustomerName(customerName)
	return sendISmartResponse(iSmart, err)
}

// GetISmartByFiberNode implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByFiberNode(fiberNode string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByFiberNode(fiberNode)
	return sendISmartResponse(iSmart, err)
}

// GetISmartByID implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByID(id string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByID(id)
	return sendISmartResponse(iSmart, err)
}

// GetISmartByStreet implements ISmartService.
func (i *iSmartServiceImpl) GetISmartByStreet(street string) (response.ISmartResponse, error) {
	iSmart, err := i.ismartRepository.GetISmartByStreet(street)
	return sendISmartResponse(iSmart, err)
}

// GetISmartsWithFilters implements ISmartService.
func (i *iSmartServiceImpl) GetISmartsWithFilters(filters map[string]interface{}) ([]response.ISmartResponse, error) {
	iSmarts, err := i.ismartRepository.GetISmartsWithFilters(filters)
	return sendISmartsResponse(iSmarts, err)
}

// UpdateISmart implements ISmartService.
func (i *iSmartServiceImpl) UpdateISmart(id string, iSmart request.ISmartRequest) (response.ISmartResponse, error) {
	// Get existing ISmart by ID
	existingISmart, err := i.ismartRepository.GetISmartByID(id)
	if err != nil {
		return response.ISmartResponse{}, err
	}

	if existingISmart.FiberNode == "" {
		sanitize := utils.SanitizeString(iSmart.FiberNode)
		existingISmart.FiberNode = sanitize
	}
	if existingISmart.Address == "" {
		sanitize := utils.SanitizeString(iSmart.Address)
		existingISmart.Address = sanitize
	}
	if existingISmart.CustomerName == "" {
		sanitize := utils.SanitizeString(iSmart.CustomerName)
		existingISmart.CustomerName = sanitize
	}
	if existingISmart.Coordinate == "" {
		sanitize := utils.SanitizeString(iSmart.Coordinate)
		existingISmart.Coordinate = sanitize
	}
	if existingISmart.Street == "" {
		sanitize := utils.SanitizeString(iSmart.Street)
		existingISmart.Street = sanitize
	}

	updatedISmart, err := i.ismartRepository.UpdateISmart(existingISmart)
	return sendISmartResponse(updatedISmart, err)
}

func NewISmartServiceImpl(ismartRepository repo.ISmartRepo) ISmartService {
	return &iSmartServiceImpl{
		ismartRepository: ismartRepository,
	}
}
