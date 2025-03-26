package services

import (
	"errors"
	"os"

	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
	"gorm.io/gorm"
)

type UserAuthServiceImpl struct {
	UserRepo repo.UserRepo
}

// Register implements UserAuth.
func (u *UserAuthServiceImpl) Register(user request.RegisterUserRequest) (response.RegisterUserResponse, error) {
	user.Name = utils.SanitizeString(user.Name)
	user.Email = utils.SanitizeString(user.Email)
	user.Phone = utils.SanitizeString(user.Phone)
	if user.CallSign != "" {
		sanitizedCallSign := utils.SanitizeString(user.CallSign)
		user.CallSign = sanitizedCallSign
	}
	if user.Contractor != nil {
		sanitizeContractor := utils.SanitizeString(*user.Contractor)
		user.Contractor = &sanitizeContractor
	}

	if !utils.ValidateEmail(user.Email) {
		return response.RegisterUserResponse{}, errors.New("invalid email")
	}

	_, findEmailErr := u.UserRepo.GetUserByEmail(user.Email)
	if findEmailErr == nil {
		return response.RegisterUserResponse{}, errors.New("user with that email already exists")
	} else if !errors.Is(findEmailErr, gorm.ErrRecordNotFound) {
		return response.RegisterUserResponse{}, findEmailErr
	}

	_, findPhoneErr := u.UserRepo.GetUserByPhone(user.Phone)
	if findPhoneErr == nil {
		return response.RegisterUserResponse{}, errors.New("user with that phone already exists")
	} else if !errors.Is(findPhoneErr, gorm.ErrRecordNotFound) {
		return response.RegisterUserResponse{}, findPhoneErr
	}
	userModel := models.User{
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Password:   user.Password,
		CallSign:   user.CallSign,
		Contractor: user.Contractor,
		Status:     func(s string) *string { return &s }("active"),
	}

	// Set default role to "user"
	role, err := u.UserRepo.GetRoleByRoleName("user")
	if err != nil {
		return response.RegisterUserResponse{}, err
	}
	userModel.Role = &role
	userModel.RoleID = &role.ID

	userCreated, err := u.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}
	return response.RegisterUserResponse{
		ID: userCreated.ID,
	}, nil
}

// Login implements UserAuth.
func (u *UserAuthServiceImpl) Login(users request.LoginUserRequest) (response.LoginUserResponse, error) {
	var user models.User
	var err error

	// Try to find user by email
	user, err = u.UserRepo.GetUserByEmail(users.Email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		user, err = u.UserRepo.GetUserByPhone(users.Phone)
		if err != nil {
			return response.LoginUserResponse{}, err
		}
	} else if err != nil {
		return response.LoginUserResponse{}, err
	}

	pass, err := utils.Decrypt(user.Password, os.Getenv("DB_KEY_ENCRYPT"))
	if err != nil {
		return response.LoginUserResponse{}, err
	}

	if pass != users.Password {
		return response.LoginUserResponse{}, errors.New("invalid credentials")
	}

	token := utils.GenerateJWTToken(user.ID)
	csrfToken := utils.GenerateCSRFToken(32)

	return response.LoginUserResponse{
		ID:           user.ID,
		SessionToken: token,
		CsrfToken:    csrfToken,
	}, nil
}

// Logout implements UserAuth.
func (u *UserAuthServiceImpl) Logout(users request.LogoutUserRequest) error {
	return nil
}

// Validate implements UserAuthService.
func (u *UserAuthServiceImpl) Validate(userId string) (response.LoginUserResponse, error) {
	user, err := u.UserRepo.GetUserById(userId)
	if err != nil {
		return response.LoginUserResponse{}, err
	}

	token := utils.GenerateJWTToken(user.ID)
	csrf := utils.GenerateCSRFToken(32)

	return response.LoginUserResponse{
		ID:           user.ID,
		SessionToken: token,
		CsrfToken:    csrf,
	}, nil
}

func NewAuthService(userRepo repo.UserRepo) UserAuthService {
	return &UserAuthServiceImpl{
		UserRepo: userRepo,
	}
}
