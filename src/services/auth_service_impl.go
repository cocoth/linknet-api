package services

import (
	"errors"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
	"gorm.io/gorm"
)

type UserAuthServiceImpl struct {
	UserRepo repo.UserRepo
}

// Register implements UserAuth.
func (u *UserAuthServiceImpl) Register(users request.RegisterUserRequest) (response.RegisterUserResponse, error) {

	_, findEmailErr := u.UserRepo.GetUserByEmail(users.Email)
	if findEmailErr == nil {
		return response.RegisterUserResponse{}, errors.New("user with that email already exists")
	} else if !errors.Is(findEmailErr, gorm.ErrRecordNotFound) {
		return response.RegisterUserResponse{}, findEmailErr
	}

	_, findPhoneErr := u.UserRepo.GetUserByPhone(users.Phone)
	if findPhoneErr == nil {
		return response.RegisterUserResponse{}, errors.New("user with that phone already exists")
	} else if !errors.Is(findPhoneErr, gorm.ErrRecordNotFound) {
		return response.RegisterUserResponse{}, findPhoneErr
	}
	userModel := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Phone:    users.Phone,
		Password: users.Password,
	}

	user, err := u.UserRepo.Create(userModel)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}
	return response.RegisterUserResponse{
		Id: user.Id,
	}, nil
}

// Login implements UserAuth.
func (u *UserAuthServiceImpl) Login(users request.LoginUserRequest) (response.LoginUserResponse, error) {
	var user models.User
	var err error

	// Try to find user by email
	user, err = u.UserRepo.GetUserByEmail(users.Email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// If not found by email, try to find user by phone
		user, err = u.UserRepo.GetUserByPhone(users.Phone)
		if err != nil {
			return response.LoginUserResponse{}, err
		}
	} else if err != nil {
		return response.LoginUserResponse{}, err
	}

	// Generate token
	errPass := utils.CompareHashPassword([]byte(users.Password), user.Password)
	if errPass != nil {
		a := "error pass" + errPass.Error()
		utils.Error(a)
		return response.LoginUserResponse{}, errors.New("invalid credentials")
	}

	// Generate token (assuming you have a function to generate a token)
	token := utils.GenerateJWTToken(user.Id)
	return response.LoginUserResponse{
		Id:    user.Id,
		Token: token,
	}, nil
}

// Logout implements UserAuth.
func (u *UserAuthServiceImpl) Logout(users request.LogoutUserRequest) error {
	panic("unimplemented")
}

// RefreshToken implements UserAuth.
func (u *UserAuthServiceImpl) RefreshToken(users request.RefreshTokenRequest) (response.RefreshTokenResponse, error) {
	panic("unimplemented")
}

func NewAuthService(userRepo repo.UserRepo) UserAuthService {
	return &UserAuthServiceImpl{
		UserRepo: userRepo,
	}
}
