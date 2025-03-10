package services

import (
	"errors"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/data/request"
	"github.com/cocoth/linknet-api/src/data/response"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UserRepo repo.UserRepo
	Validate *validator.Validate
}

// Create implements UserService.
func (u *UsersServiceImpl) Create(users request.CreateUserReq) (response.UserRes, error) {
	err := u.Validate.Struct(users)
	if err != nil {
		return response.UserRes{}, err
	}

	existingUser, err := u.UserRepo.FindUserByEmail(users.Email)
	if err == nil && existingUser.Id != "" {
		return response.UserRes{}, errors.New("email already exists")
	}

	userModel := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
		// Role: models.Role{
		// 	Name: users.Role.Name,
		// },
	}
	var role models.Role
	if err := u.UserRepo.GetOrCreateRole(users.Role.Name, &role); err != nil {
		utils.ErrPanic(err)
		return response.UserRes{}, err

	}
	userModel.Role = role
	if _, err := u.UserRepo.Create(userModel); err != nil {
		return response.UserRes{}, err
	}
	return response.UserRes{
		Id:    userModel.Id,
		Name:  userModel.Name,
		Email: userModel.Email,
		Role:  userModel.Role.Name,
	}, nil
}

// Delete implements UserService.
func (u *UsersServiceImpl) Delete(id string) error {
	u.UserRepo.Delete(id)
	return nil
}

// GetAll implements UserService.
func (u *UsersServiceImpl) GetAll() ([]response.UserRes, error) {
	result := u.UserRepo.GetAll()
	var users []response.UserRes
	for _, user := range result {
		users = append(users, response.UserRes{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role.Name,
		})
	}

	return users, nil
}

// GetById implements UserService.
func (u *UsersServiceImpl) GetById(id string) (response.UserRes, error) {
	user := u.UserRepo.GetById(id)
	return response.UserRes{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Update implements UserService.
func (u *UsersServiceImpl) Update(users request.UpdateUserReq) (response.UserRes, error) {
	user := u.UserRepo.GetById(users.Id)
	user.Name = users.Name
	user.Email = users.Email
	u.UserRepo.Update(user)
	return response.UserRes{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role.Name,
	}, nil

}

func NewUserServiceImpl(user repo.UserRepo, validate *validator.Validate) UserService {
	return &UsersServiceImpl{
		UserRepo: user,
		Validate: validate,
	}
}
