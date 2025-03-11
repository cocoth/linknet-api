package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
)

type UserService interface {
	GetAll() ([]response.UserResponse, error)
	GetById(id string) (response.UserResponse, error)
	Create(users request.CreateUserRequest) (response.UserResponse, error)
	Update(users request.UpdateUserRequest) (response.UserResponse, error)
	Delete(id string) error
	// TODO!
	// GetByEmail(users request.GetUserByEmailReq) (request.CreateUserReq, error)
	// GetByPhone(users request.GetUserByPhoneReq) (request.CreateUserReq, error)
	// GetByRole(users request.GetUserByRoleReq) ([]request.CreateUserReq, error)
	// GetByStatus(users request.GetUserByStatusReq) ([]request.CreateUserReq, error)
}

type UserAuthService interface {
	Register(users request.RegisterUserRequest) (response.RegisterUserResponse, error)
	Login(users request.LoginUserRequest) (response.LoginUserResponse, error)
	Logout(users request.LogoutUserRequest) error
	RefreshToken(users request.RefreshTokenRequest) (response.RefreshTokenResponse, error)
}
