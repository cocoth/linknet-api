package repo

import (
	"errors"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/http/request"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	Db *gorm.DB
}

// GetOrCreateRole implements UserRepo.
func (u *userRepoImpl) GetOrCreateRole(name string, role *models.Role) error {
	err := u.Db.Where("name = ?", name).First(role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newRole := models.Role{Name: name}
		if err := u.Db.Create(&newRole).Error; err != nil {
			return err
		}
		*role = newRole
	} else if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail implements UserRepo.
func (u *userRepoImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserByPhone implements UserRepo.
func (u *userRepoImpl) GetUserByPhone(phone string) (models.User, error) {
	var user models.User
	err := u.Db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Create implements UserRepo.
func (u *userRepoImpl) Create(user models.User) (models.User, error) {
	result := u.Db.Create(&user)
	return user, result.Error
}

// Delete implements UserRepo.
func (u *userRepoImpl) Delete(id string) error {
	result := u.Db.Delete(&models.User{}, "id?=", id)
	return result.Error
}

// GetAll implements UserRepo.
func (u *userRepoImpl) GetAll() []models.User {

	var users []models.User
	result := u.Db.Preload("Role").Find(&users)
	if result.Error != nil {
		return nil
	}
	return users
}

// GetById implements UserRepo.
func (u *userRepoImpl) GetById(id string) (models.User, error) {
	var user models.User

	result := u.Db.Where("id = ?", id).First(&user)
	return user, result.Error
}

// Update implements UserRepo.
func (u *userRepoImpl) Update(user models.User) (models.User, error) {
	var updateUser = request.UpdateUserRequest{
		Id:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		RoleID:   *user.RoleID,
	}

	result := u.Db.Model(&user).Where("id = ?", user.Id).Updates(updateUser)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil

}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &userRepoImpl{Db: db}

}
