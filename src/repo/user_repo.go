package repo

import "github.com/cocoth/linknet-api/config/models"

type UserRepo interface {
	GetOrCreateRole(name string, role *models.Role) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)

	GetAll() []models.User
	GetById(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string) error
}
