package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UsersServiceImpl struct {
	UserRepo repo.UserRepo
	Validate *validator.Validate
}

func sendUserResponse(userModel models.User, err error) (response.UserResponse, error) {
	if err != nil {
		return response.UserResponse{}, err
	}
	var roleName string
	if userModel.Role != nil {
		roleName = userModel.Role.Name
	}
	pass, err := utils.Decrypt(userModel.Password, os.Getenv("DB_KEY_ENCRYPT"))
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.UserResponse{
		ID:         userModel.ID,
		Name:       userModel.Name,
		Email:      userModel.Email,
		Phone:      userModel.Phone,
		Password:   &pass,
		CallSign:   userModel.CallSign,
		Contractor: userModel.Contractor,
		Status:     userModel.Status,
		Role: &response.RoleResponse{
			ID:   strconv.FormatUint(uint64(*userModel.RoleID), 10),
			Name: roleName,
		},
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
		DeletedAt: userModel.DeletedAt,
	}, nil
}

func sendUsersResponse(userModel []models.User, err error) ([]response.UserResponse, error) {
	if err != nil {
		return nil, err
	}
	users := make([]response.UserResponse, 0, len(userModel))
	for _, user := range userModel {
		var roleResp *response.RoleResponse
		if user.Role != nil {
			roleResp = &response.RoleResponse{
				ID:   strconv.FormatUint(uint64(*user.RoleID), 10),
				Name: user.Role.Name,
			}
		}
		pass, err := utils.Decrypt(user.Password, os.Getenv("DB_KEY_ENCRYPT"))
		if err != nil {
			return nil, err
		}
		users = append(users, response.UserResponse{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Phone:      user.Phone,
			Password:   &pass,
			CallSign:   user.CallSign,
			Contractor: user.Contractor,
			Status:     user.Status,
			Role:       roleResp,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			DeletedAt:  user.DeletedAt,
		})
	}
	return users, nil
}

// GetAdmins implements UserService.
func (u *UsersServiceImpl) GetAdmins() ([]response.UserResponse, error) {
	result, err := u.UserRepo.GetAdmins()
	return sendUsersResponse(result, err)
}

// GetUsersWithFilters implements UserService.
func (u *UsersServiceImpl) GetUsersWithFilters(filters map[string]interface{}) ([]response.UserResponse, error) {
	result, err := u.UserRepo.GetUsersWithFilters(filters)
	return sendUsersResponse(result, err)
}

// GetUsersByCallSign implements UserService.
func (u *UsersServiceImpl) GetUsersByCallSign(callsign string) ([]response.UserResponse, error) {
	result, err := u.UserRepo.GetUsersByCallSign(callsign)
	return sendUsersResponse(result, err)
}

// GetAll implements UserService.
func (u *UsersServiceImpl) GetAll() ([]response.UserResponse, error) {
	result := u.UserRepo.GetAll()
	return sendUsersResponse(result, nil)

}

// GetUserById implements UserService.
func (u *UsersServiceImpl) GetUserById(id string) (response.UserResponse, error) {
	user, err := u.UserRepo.GetUserById(id)
	return sendUserResponse(user, err)
}

// GetUsersByEmail implements UserService.
func (u *UsersServiceImpl) GetUsersByEmail(email string) ([]response.UserResponse, error) {

	err := u.Validate.Var(email, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByEmail(email)

	return sendUsersResponse(dataUsers, err)
}

// GetUsersByName implements UserService.
func (u *UsersServiceImpl) GetUsersByName(name string) ([]response.UserResponse, error) {

	err := u.Validate.Var(name, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByName(name)

	return sendUsersResponse(dataUsers, err)
}

// GetUsersByPhone implements UserService.
func (u *UsersServiceImpl) GetUsersByPhone(phone string) ([]response.UserResponse, error) {

	err := u.Validate.Var(phone, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByPhone(phone)
	return sendUsersResponse(dataUsers, err)
}

// GetUsersByRole implements UserService.
func (u *UsersServiceImpl) GetUsersByRole(role string) ([]response.UserResponse, error) {

	err := u.Validate.Var(role, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByRole(role)

	return sendUsersResponse(dataUsers, err)
}

// GetUsersByStatus implements UserService.
func (u *UsersServiceImpl) GetUsersByStatus(status string) ([]response.UserResponse, error) {

	err := u.Validate.Var(status, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByStatus(status)
	return sendUsersResponse(dataUsers, err)
}

// GetUsersByContractor implements UserService.
func (u *UsersServiceImpl) GetUsersByContractor(contractor string) ([]response.UserResponse, error) {

	err := u.Validate.Var(contractor, "required")
	if err != nil {
		return nil, err
	}
	dataUsers, err := u.UserRepo.GetUsersByContractor(contractor)

	return sendUsersResponse(dataUsers, err)
}

// CreateUser implements UserService.
func (u *UsersServiceImpl) CreateUser(user request.UserRequest) (response.UserResponse, error) {
	err := u.Validate.Struct(user)
	if err != nil {
		return response.UserResponse{}, err
	}

	user.Name = utils.SanitizeString(user.Name)
	user.Email = utils.SanitizeString(user.Email)
	user.Phone = utils.SanitizeString(user.Phone)
	if user.CallSign != "" {
		sanitizedCallSign := utils.SanitizeString(user.CallSign)
		user.CallSign = sanitizedCallSign
	}
	if user.Contractor != nil {
		sanitizeContractor := utils.SanitizeString(*user.Contractor)
		user.CallSign = sanitizeContractor
	}

	if !utils.ValidateEmail(user.Email) {
		return response.UserResponse{}, errors.New("invalid email")
	}
	if user.CallSign != "" {
		sanitizedCallSign := utils.SanitizeString(user.CallSign)
		user.CallSign = sanitizedCallSign
	}

	_, findEmailErr := u.UserRepo.GetUserByEmail(user.Email)
	if findEmailErr == nil {
		return response.UserResponse{}, errors.New("user with that email already exists")
	} else if !errors.Is(findEmailErr, gorm.ErrRecordNotFound) {
		return response.UserResponse{}, findEmailErr
	}

	_, findPhoneErr := u.UserRepo.GetUserByPhone(user.Phone)
	if findPhoneErr == nil {
		return response.UserResponse{}, errors.New("user with that phone already exists")
	} else if !errors.Is(findPhoneErr, gorm.ErrRecordNotFound) {
		return response.UserResponse{}, findPhoneErr
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

	if user.Role != nil && user.Role.Name != "" {
		role, err := u.UserRepo.GetRoleByRoleName(user.Role.Name)

		if err != nil {
			return response.UserResponse{}, err
		}
		userModel.Role = &role
		userModel.RoleID = &role.ID
	}

	createdUser, err := u.UserRepo.CreateUser(userModel)
	if err != nil {
		return response.UserResponse{}, err
	}

	return sendUserResponse(createdUser, nil)
}

// DeleteUser implements UserService.
func (u *UsersServiceImpl) DeleteUser(id string) (response.UserResponse, error) {
	user, err := u.UserRepo.DeleteUser(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	return sendUserResponse(user, nil)
}

// IsAdmin implements UserService.
func (u *UsersServiceImpl) IsAdmin(token string) (status bool, userResponse response.UserResponse, err error) {
	var user response.UserResponse

	exp, userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		return false, user, errors.New("invalid Token")
	}
	if float64(time.Now().Unix()) > exp {
		return false, user, errors.New("token expired")
	}
	userResponse, errNotAdmin := u.GetUserById(userId)
	if errNotAdmin != nil {
		return false, user, errors.New("unauthorized")
	}
	if userResponse.Role.Name != "admin" {
		return false, user, errors.New("only admin can this resource")
	}
	return true, userResponse, nil
}

// CheckToken implements UserService.
func (u *UsersServiceImpl) CheckToken(token string) (status bool, userResponse response.UserResponse, err error) {
	var user response.UserResponse
	exp, userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		return false, user, errors.New("invalid Token")
	}
	if float64(time.Now().Unix()) > exp {
		return false, user, errors.New("token expired")
	}
	userOnDB, err := u.GetUserById(userId)
	if err != nil {
		return false, user, errors.New("unauthorized")
	}

	user.ID = userOnDB.ID
	user.Name = userOnDB.Name
	user.Email = userOnDB.Email
	user.Phone = userOnDB.Phone
	user.Password = userOnDB.Password
	user.CallSign = userOnDB.CallSign
	user.Contractor = userOnDB.Contractor
	user.Status = userOnDB.Status
	user.Role = userOnDB.Role
	user.CreatedAt = userOnDB.CreatedAt
	user.UpdatedAt = userOnDB.UpdatedAt
	user.DeletedAt = userOnDB.DeletedAt

	return true, user, nil
}

// CreateRole implements UserService.
func (u *UsersServiceImpl) CreateRole(role request.RoleRequest) (response.RoleResponse, error) {
	var newRole models.Role
	err := u.UserRepo.GetOrCreateRole(role.Name, &newRole)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: role.Name}, nil
}

// GetAllRole implements UserService.
func (u *UsersServiceImpl) GetAllRole() ([]response.RoleResponse, error) {
	var roles []response.RoleResponse

	result, err := u.UserRepo.GetAllRole()
	if err != nil {
		return nil, err
	}
	for _, role := range result {
		roles = append(roles, response.RoleResponse{Name: role.Name})
	}
	return roles, nil
}

// GetRoleByRoleID implements UserService.
func (u *UsersServiceImpl) GetRoleByRoleID(roleID uint) (response.RoleResponse, error) {
	role, err := u.UserRepo.GetRoleByRoleID(roleID)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: role.Name}, nil
}

// GetRoleByRoleName implements UserService.
func (u *UsersServiceImpl) GetRoleByRoleName(roleName string) (response.RoleResponse, error) {
	role, err := u.UserRepo.GetRoleByRoleName(roleName)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: role.Name}, nil
}

// UpdateRole implements UserService.
func (u *UsersServiceImpl) UpdateRole(id uint, roleReq request.RoleRequest) (response.RoleResponse, error) {
	var role models.Role
	var err error

	role, err = u.UserRepo.GetRoleByRoleID(id)

	if err != nil {
		return response.RoleResponse{}, err
	}

	updatedRole, err := u.UserRepo.UpdateRole(role)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: updatedRole.Name}, nil
}

// DeleteRoleByID implements UserService.
func (u *UsersServiceImpl) DeleteRoleByID(roleID uint) (response.RoleResponse, error) {
	role, err := u.UserRepo.DeleteRoleByID(roleID)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: role.Name}, nil
}

// DeleteRoleByName implements UserService.
func (u *UsersServiceImpl) DeleteRoleByName(roleName string) (response.RoleResponse, error) {
	role, err := u.UserRepo.DeleteRoleByName(roleName)
	if err != nil {
		return response.RoleResponse{}, err
	}
	return response.RoleResponse{Name: role.Name}, nil
}

// UpdateUser implements UserService.
func (u *UsersServiceImpl) UpdateUser(id string, user request.UpdateUserRequest) (response.UserResponse, error) {
	userData, err := u.UserRepo.GetUserById(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	if user.Name != nil {
		sanitizedName := utils.SanitizeString(*user.Name)
		user.Name = &sanitizedName
	}
	if user.Email != nil {
		sanitizedEmail := utils.SanitizeString(*user.Email)
		user.Email = &sanitizedEmail
	}
	if user.Phone != nil {
		sanitizedPhone := utils.SanitizeString(*user.Phone)
		user.Phone = &sanitizedPhone
	}
	if user.CallSign != nil {
		sanitizedCallSign := utils.SanitizeString(*user.CallSign)
		user.CallSign = &sanitizedCallSign
	}
	if user.Contractor != nil {
		sanitizedContractor := utils.SanitizeString(*user.Contractor)
		user.Contractor = &sanitizedContractor
	}

	if user.Email != nil && !utils.ValidateEmail(*user.Email) {
		return response.UserResponse{}, errors.New("invalid email")
	}

	if user.Name != nil {
		userData.Name = *user.Name
	}
	if user.Email != nil {
		userData.Email = *user.Email
	}

	if user.Phone != nil {
		userData.Phone = *user.Phone
	}
	if user.Password != nil {
		hash, err := utils.Encrypt(*user.Password, os.Getenv("DB_KEY_ENCRYPT"))
		if err != nil {
			return response.UserResponse{}, err
		}
		userData.Password = hash
	}
	if user.CallSign != nil {
		userData.CallSign = *user.CallSign
	}
	if user.Contractor != nil {
		userData.Contractor = user.Contractor
	}
	if user.Status != nil {
		userData.Status = user.Status
	}

	if user.Role != nil && user.Role.Name != "" {
		role, err := u.UserRepo.GetRoleByRoleName(user.Role.Name)
		if err != nil {
			return response.UserResponse{}, err
		}
		userData.Role = &role
		userData.RoleID = &role.ID
	}
	updatedUser, err := u.UserRepo.UpdateUser(userData)
	if err != nil {
		return response.UserResponse{}, err
	}

	return sendUserResponse(updatedUser, err)
}

func NewUserServiceImpl(user repo.UserRepo, validate *validator.Validate) UserService {
	return &UsersServiceImpl{
		UserRepo: user,
		Validate: validate,
	}
}
