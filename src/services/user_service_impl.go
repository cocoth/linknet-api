package services

import (
	"errors"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UsersServiceImpl struct {
	UserRepo repo.UserRepo
	Validate *validator.Validate
}

func sendUserResponseponse(userModel models.User, err error) (response.UserResponse, error) {
	if err != nil {
		return response.UserResponse{}, err
	}
	var roleName string
	if userModel.Role != nil {
		roleName = userModel.Role.Name
	}
	return response.UserResponse{
		Id:    userModel.Id,
		Name:  userModel.Name,
		Email: userModel.Email,
		Phone: userModel.Phone,
		Role:  &response.RoleResponse{Name: roleName},
	}, nil
}

// Create implements UserService.
func (u *UsersServiceImpl) Create(users request.CreateUserRequest) (response.UserResponse, error) {
	err := u.Validate.Struct(users)
	if err != nil {
		return response.UserResponse{}, err
	}

	_, findEmailErr := u.UserRepo.GetUserByEmail(users.Email)
	if findEmailErr == nil {
		return response.UserResponse{}, errors.New("user with that email already exists")
	} else if !errors.Is(findEmailErr, gorm.ErrRecordNotFound) {
		return response.UserResponse{}, findEmailErr
	}

	_, findPhoneErr := u.UserRepo.GetUserByPhone(users.Phone)
	if findPhoneErr == nil {
		return response.UserResponse{}, errors.New("user with that phone already exists")
	} else if !errors.Is(findPhoneErr, gorm.ErrRecordNotFound) {
		return response.UserResponse{}, findPhoneErr
	}

	userModel := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Phone:    users.Phone,
		Password: users.Password,
		Role: &models.Role{
			Name: users.Role.Name,
		},
	}

	var role models.Role
	if err := u.UserRepo.GetOrCreateRole(users.Role.Name, &role); err != nil {
		return response.UserResponse{}, err
	}

	userModel.RoleID = &role.ID

	user, err := u.UserRepo.Create(userModel)
	if err != nil {
		return response.UserResponse{}, err
	}

	return sendUserResponseponse(user, nil)
}

// Delete implements UserService.
func (u *UsersServiceImpl) Delete(id string) error {
	err := u.UserRepo.Delete(id)
	return err
}

// GetAll implements UserService.
func (u *UsersServiceImpl) GetAll() ([]response.UserResponse, error) {
	result := u.UserRepo.GetAll()
	var users []response.UserResponse

	for _, user := range result {
		var roleResp *response.RoleResponse
		if user.Role != nil {
			roleResp = &response.RoleResponse{
				Name: user.Role.Name,
			}
		}
		users = append(users, response.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Role:  roleResp,
		})
	}

	return users, nil
}

// GetById implements UserService.
func (u *UsersServiceImpl) GetById(id string) (response.UserResponse, error) {
	user, err := u.UserRepo.GetById(id)
	return sendUserResponseponse(user, err)
}

// Update implements UserService.
func (u *UsersServiceImpl) Update(users request.UpdateUserRequest) (response.UserResponse, error) {
	user, err := u.UserRepo.GetById(users.Id)
	user.Name = users.Name
	user.Email = users.Email
	u.UserRepo.Update(user)
	return sendUserResponseponse(user, err)

}

func NewUserServiceImpl(user repo.UserRepo, validate *validator.Validate) UserService {
	return &UsersServiceImpl{
		UserRepo: user,
		Validate: validate,
	}
}
