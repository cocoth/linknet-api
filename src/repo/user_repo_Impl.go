package repo

import (
	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/data/request"
	"github.com/cocoth/linknet-api/src/utils"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	Db *gorm.DB
}

// GetOrCreateRole implements UserRepo.
func (u *userRepoImpl) GetOrCreateRole(name string, role *models.Role) error {
	if err := u.Db.Where("name = ?", name).First(role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			role.Name = name
			if err := u.Db.Create(role).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// FindUserByEmail implements UserRepo.
func (u *userRepoImpl) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.Db.Where("email = ?", email).First(&user).Error
	return user, err
}

// Create implements UserRepo.
func (u *userRepoImpl) Create(user models.User) (models.User, error) {
	result := u.Db.Create(&user)
	utils.ErrPanic(result.Error)
	return user, nil
}

// Delete implements UserRepo.
func (u *userRepoImpl) Delete(id string) {
	var user models.User

	result := u.Db.Where("id = ?", id).Delete(&user)
	utils.ErrPanic(result.Error)
}

// GetAll implements UserRepo.
func (u *userRepoImpl) GetAll() []models.User {

	var users []models.User
	result := u.Db.Find(&users)
	utils.ErrPanic(result.Error)

	return users
}

// GetById implements UserRepo.
func (u *userRepoImpl) GetById(id string) models.User {
	var user models.User

	result := u.Db.Where("id = ?", id).First(&user)
	utils.ErrPanic(result.Error)
	return user
}

// Update implements UserRepo.
func (u *userRepoImpl) Update(user models.User) (models.User, error) {
	var updateUser = request.UpdateUserReq{
		Id:       user.Id,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	result := u.Db.Model(&user).Where("id = ?", user.Id).Updates(updateUser)
	utils.ErrPanic(result.Error)
	return user, nil

}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &userRepoImpl{Db: db}

}
