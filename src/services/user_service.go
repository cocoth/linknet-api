package services

import (
	"github.com/cocoth/linknet-api/src/data/request"
	"github.com/cocoth/linknet-api/src/data/response"
)

type UserService interface {
	GetAll() ([]response.UserRes, error)
	GetById(id string) (response.UserRes, error)
	Create(users request.CreateUserReq) (response.UserRes, error)
	Update(users request.UpdateUserReq) (response.UserRes, error)
	Delete(id string) error

	// TODO!
	// GetByEmail(users request.GetUserByEmailReq) (request.CreateUserReq, error)
	// GetByPhone(users request.GetUserByPhoneReq) (request.CreateUserReq, error)
	// GetByRole(users request.GetUserByRoleReq) ([]request.CreateUserReq, error)
	// GetByStatus(users request.GetUserByStatusReq) ([]request.CreateUserReq, error)
}
