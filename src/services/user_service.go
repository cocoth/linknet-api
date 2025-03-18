package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
)

type UserService interface {
	GetUsersByName(name string) ([]response.UserResponse, error)
	GetUsersByEmail(email string) ([]response.UserResponse, error)
	GetUsersByPhone(phone string) ([]response.UserResponse, error)
	GetUsersByRole(role string) ([]response.UserResponse, error)
	GetUsersByStatus(status string) ([]response.UserResponse, error)
	GetUsersByContractor(contractor string) ([]response.UserResponse, error)

	IsAdmin(token string) (bool, error)
	GetAllRole() ([]response.RoleResponse, error)
	CreateRole(role request.RoleRequest) (response.RoleResponse, error)
	GetAll() ([]response.UserResponse, error)
	GetById(id string) (response.UserResponse, error)
	Create(user request.UserRequest) (response.UserResponse, error)
	Update(id string, user request.UpdateUserRequest) (response.UserResponse, error)
	Delete(id string) (response.UserResponse, error)
}

type UserAuthService interface {
	Register(user request.RegisterUserRequest) (response.RegisterUserResponse, error)
	Login(users request.LoginUserRequest) (response.LoginUserResponse, error)
	Logout(users request.LogoutUserRequest) error
	Validate(token string) (response.LoginUserResponse, error)
}
