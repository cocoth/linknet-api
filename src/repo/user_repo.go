package repo

import "github.com/cocoth/linknet-api/config/models"

type UserRepo interface {
	GetAll() []models.User
	GetById(id string) models.User
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string)
	GetOrCreateRole(name string, role *models.Role) error
	FindUserByEmail(email string) (models.User, error)
	// FindUserByPhone(phone string) (models.User, error)
}
