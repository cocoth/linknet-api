package repo

import "github.com/cocoth/linknet-api/config/models"

type UserRepo interface {
	GetAllRole() ([]models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	GetOrCreateRole(name string, role *models.Role) error

	GetRoleByRoleName(role string) (models.Role, error)
	GetRoleByUserName(name string) (models.Role, error)
	GetRoleByUserPhone(phone string) (models.Role, error)
	GetRoleByUserEmail(email string) (models.Role, error)
	GetUserByContractor(contractor string) (models.User, error)

	GetDeletedUserByEmail(email string) (models.User, error)
	GetDeletedUserByName(name string) (models.User, error)
	GetDeletedUserByPhone(phone string) (models.User, error)

	GetUsersByName(name string) ([]models.User, error)
	GetUsersByEmail(email string) ([]models.User, error)
	GetUsersByPhone(phone string) ([]models.User, error)
	GetUsersByRole(role string) ([]models.User, error)
	GetUsersByStatus(status string) ([]models.User, error)
	GetUsersByContractor(contractor string) ([]models.User, error)

	GetDeletedUsersByEmail(email string) ([]models.User, error)
	GetDeletedUsersByName(name string) ([]models.User, error)
	GetDeletedUsersByPhone(phone string) ([]models.User, error)

	GetUserByName(name string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)

	GetAll() []models.User
	GetById(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string) (models.User, error)
}
