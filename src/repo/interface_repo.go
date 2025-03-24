package repo

import "github.com/cocoth/linknet-api/src/models"

type UserRepo interface {
	GetAllRole() ([]models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	GetOrCreateRole(name string, role *models.Role) error
	UpdateRole(role models.Role) (models.Role, error)
	DeleteRoleByID(roleID uint) (models.Role, error)
	DeleteRoleByName(roleName string) (models.Role, error)

	GetRoleByRoleID(roleID uint) (models.Role, error)
	GetRoleByRoleName(role string) (models.Role, error)
	GetRoleByUserName(name string) (models.Role, error)
	GetRoleByUserPhone(phone string) (models.Role, error)
	GetRoleByUserEmail(email string) (models.Role, error)

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

	GetAll() []models.User
	GetUserById(id string) (models.User, error)
	GetUserByName(name string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByPhone(phone string) (models.User, error)
	GetUserByContractor(contractor string) (models.User, error)

	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
}

type FileUploadRepo interface {
	UploadFile(file models.FileUpload) (models.FileUpload, error)
	GetAllFileUpload() ([]models.FileUpload, error)

	GetFileUploadByFileID(id string) (models.FileUpload, error)
	GetFileUploadByFileName(fileName string) (models.FileUpload, error)
	GetFileUploadByFileHash(fileHash string) (models.FileUpload, error)

	GetFilesUploadByAuthorID(authorID string) ([]models.FileUpload, error)
	GetFilesUploadByAuthorName(authorName string) ([]models.FileUpload, error)
	GetFilesUploadByFileName(fileName string) ([]models.FileUpload, error)

	UpdateFileUpload(file models.FileUpload) (models.FileUpload, error)

	DeleteFileUploadByFileID(id string) (models.FileUpload, error)
	DeleteFileUploadByFileName(fileName string) (models.FileUpload, error)
	DeleteFileUploadByFileHash(fileHash string) (models.FileUpload, error)
}
