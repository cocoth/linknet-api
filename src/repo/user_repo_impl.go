package repo

import (
	"errors"

	"github.com/cocoth/linknet-api/src/models"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	Db *gorm.DB
}

// GetAllRole implements UserRepo.
func (u *userRepoImpl) GetAllRole() ([]models.Role, error) {
	var roles []models.Role
	err := u.Db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// CreateRole implements UserRepo.
func (u *userRepoImpl) CreateRole(role models.Role) (models.Role, error) {
	err := u.Db.Create(&role).Error
	return role, err
}

// GetOrCreateRole implements UserRepo.
func (u *userRepoImpl) GetOrCreateRole(name string, role *models.Role) error {
	existingRole, err := u.GetRoleByRoleName(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newRole := models.Role{Name: name}
		createdRole, err := u.CreateRole(newRole)
		if err != nil {
			return err
		}
		*role = createdRole
	} else if err != nil {
		return err
	} else {
		*role = existingRole
	}
	return nil
}

// UpdateRole implements UserRepo.
func (u *userRepoImpl) UpdateRole(role models.Role) (models.Role, error) {
	updatedRole := map[string]interface{}{}

	if role.ID != 0 {
		updatedRole["id"] = role.ID
	} else if role.Name != "" {
		updatedRole["name"] = role.Name
	} else {
		return role, nil
	}

	if len(updatedRole) == 0 {
		return role, nil
	}

	result := u.Db.Model(&role).Where("id = ?", role.ID).Updates(updatedRole)
	if result.Error != nil {
		return role, result.Error
	}

	err := u.Db.Preload("Role").First(&role, "id = ?", role.ID).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

// DeleteRoleByID implements UserRepo.
func (u *userRepoImpl) DeleteRoleByID(roleID uint) (models.Role, error) {
	var role models.Role
	err := u.Db.Where("id = ?", roleID).Delete(&role).Error
	return role, err
}

// DeleteRoleByName implements UserRepo.
func (u *userRepoImpl) DeleteRoleByName(roleName string) (models.Role, error) {
	var role models.Role
	err := u.Db.Where("name = ?", roleName).Delete(&role).Error
	return role, err
}

// GetRoleByRoleID implements UserRepo.
func (u *userRepoImpl) GetRoleByRoleID(roleID uint) (models.Role, error) {
	var roleName models.Role
	err := u.Db.Where("id = ?", roleID).First(&roleName).Error
	return roleName, err
}

// GetRoleByRoleName implements UserRepo.
func (u *userRepoImpl) GetRoleByRoleName(role string) (models.Role, error) {
	var roleName models.Role
	err := u.Db.Where("name = ?", role).First(&roleName).Error
	return roleName, err
}

// GetRoleByName implements UserRepo.
func (u *userRepoImpl) GetRoleByUserName(name string) (models.Role, error) {
	var role models.Role
	err := u.Db.Joins("JOIN users ON users.role_id = roles.id").Where("users.name = ?", name).First(&role).Error

	return role, err
}

// GetRoleByEmail implements UserRepo.
func (u *userRepoImpl) GetRoleByUserEmail(email string) (models.Role, error) {
	var role models.Role
	err := u.Db.Joins("JOIN users ON users.role_id = roles.id").Where("users.email = ?", email).First(&role).Error
	return role, err
}

// GetRoleByPhone implements UserRepo.
func (u *userRepoImpl) GetRoleByUserPhone(phone string) (models.Role, error) {
	var role models.Role
	err := u.Db.Joins("JOIN users ON users.role_id = roles.id").Where("users.phone = ?", phone).First(&role).Error
	return role, err
}

// GetUserByName implements UserRepo.
func (u *userRepoImpl) GetUsersByName(name string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail implements UserRepo.
func (u *userRepoImpl) GetUsersByEmail(email string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Where("email LIKE ?", "%"+email+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByPhone implements UserRepo.
func (u *userRepoImpl) GetUsersByPhone(phone string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Where("phone LIKE ?", "%"+phone+"%").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByRole implements UserRepo.
func (u *userRepoImpl) GetUsersByRole(role string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Joins("JOIN roles ON roles.id = users.role_id").Where("roles.name = ?", role).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByStatus implements UserRepo.
func (u *userRepoImpl) GetUsersByStatus(status string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Where("status = ?", status).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByContractor implements UserRepo.
func (u *userRepoImpl) GetUsersByContractor(contractor string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Preload("Role").Where("contractor = ?", contractor).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetDeletedUserByEmail implements UserRepo.
func (u *userRepoImpl) GetDeletedUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.Db.Unscoped().Where("email = ?", email).Find(&user).Error
	return user, err
}

// GetDeletedUserByName implements UserRepo.
func (u *userRepoImpl) GetDeletedUserByName(name string) (models.User, error) {
	var user models.User
	err := u.Db.Unscoped().Where("name = ?", name).Find(&user).Error
	return user, err
}

// GetDeletedUserByPhone implements UserRepo.
func (u *userRepoImpl) GetDeletedUserByPhone(phone string) (models.User, error) {
	var user models.User
	err := u.Db.Unscoped().Where("phone = ?", phone).Find(&user).Error
	return user, err
}

// GetUserByEmail implements UserRepo.
func (u *userRepoImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.Db.Preload("Role").First(&user, "email = ?", email).Error
	return user, err
}

// GetUserByName implements UserRepo.
func (u *userRepoImpl) GetUserByName(name string) (models.User, error) {
	var user models.User
	err := u.Db.Preload("Role").First(&user, "name = ?", name).Error
	return user, err
}

// GetUserByPhone implements UserRepo.
func (u *userRepoImpl) GetUserByPhone(phone string) (models.User, error) {
	var user models.User
	err := u.Db.Preload("Role").First(&user, "phone = ?", phone).Error
	return user, err
}

// GetUserByContractor implements UserRepo.
func (u *userRepoImpl) GetUserByContractor(contractor string) (models.User, error) {
	var user models.User
	err := u.Db.Preload("Role").First(&user, "contractor = ?", contractor).Error
	return user, err
}

// GetDeletedUsersByEmail implements UserRepo.
func (u *userRepoImpl) GetDeletedUsersByEmail(email string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Unscoped().Where("email = ?", email).Find(&users).Error
	return users, err
}

// GetDeletedUsersByName implements UserRepo.
func (u *userRepoImpl) GetDeletedUsersByName(name string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Unscoped().Where("name = ?", name).Find(&users).Error
	return users, err
}

// GetDeletedUsersByPhone implements UserRepo.
func (u *userRepoImpl) GetDeletedUsersByPhone(phone string) ([]models.User, error) {
	var users []models.User
	err := u.Db.Unscoped().Where("phone = ?", phone).Find(&users).Error
	return users, err
}

// CreateUser implements UserRepo.
func (u *userRepoImpl) CreateUser(user models.User) (models.User, error) {
	result := u.Db.Create(&user)
	return user, result.Error
}

// DeleteUser implements UserRepo.
func (u *userRepoImpl) DeleteUser(id string) (models.User, error) {
	var user models.User
	if err := u.Db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	if err := u.Db.Delete(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
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

// GetUserById implements UserRepo.
func (u *userRepoImpl) GetUserById(id string) (models.User, error) {
	var user models.User

	result := u.Db.Preload("Role").Where("id = ?", id).First(&user)
	return user, result.Error
}

// UpdateUser implements UserRepo.
func (u *userRepoImpl) UpdateUser(user models.User) (models.User, error) {
	updateUser := map[string]interface{}{}

	if user.Name != "" {
		updateUser["name"] = user.Name
	}
	if user.Email != "" {
		updateUser["email"] = user.Email
	}
	if user.Phone != "" {
		updateUser["phone"] = user.Phone
	}
	if user.Password != "" {
		updateUser["password"] = user.Password
	}
	if user.CallSign != "" {
		updateUser["call_sign"] = user.CallSign
	}
	if user.Contractor != nil {
		updateUser["contractor"] = user.Contractor
	}
	if user.Status != nil {
		updateUser["status"] = user.Status
	}
	if user.Role != nil && user.Role.ID != 0 {
		updateUser["role_id"] = user.Role.ID
	}

	if len(updateUser) == 0 {
		return user, nil
	}

	result := u.Db.Model(&user).Where("id = ?", user.ID).Updates(updateUser)
	if result.Error != nil {
		return user, result.Error
	}

	err := u.Db.Preload("Role").First(&user, "id = ?", user.ID).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &userRepoImpl{Db: db}

}
