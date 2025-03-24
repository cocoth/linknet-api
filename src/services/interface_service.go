package services

import (
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
)

type UserService interface {
	GetAll() ([]response.UserResponse, error)
	GetUserById(id string) (response.UserResponse, error)

	GetUsersByName(name string) ([]response.UserResponse, error)
	GetUsersByEmail(email string) ([]response.UserResponse, error)
	GetUsersByPhone(phone string) ([]response.UserResponse, error)
	GetUsersByRole(role string) ([]response.UserResponse, error)
	GetUsersByStatus(status string) ([]response.UserResponse, error)
	GetUsersByContractor(contractor string) ([]response.UserResponse, error)

	IsAdmin(token string) (bool, error)

	CreateRole(role request.RoleRequest) (response.RoleResponse, error)
	GetAllRole() ([]response.RoleResponse, error)
	GetRoleByRoleID(roleID uint) (response.RoleResponse, error)
	GetRoleByRoleName(roleName string) (response.RoleResponse, error)
	UpdateRole(id uint, roleReq request.RoleRequest) (response.RoleResponse, error)
	DeleteRoleByID(roleID uint) (response.RoleResponse, error)
	DeleteRoleByName(roleName string) (response.RoleResponse, error)

	CreateUser(user request.UserRequest) (response.UserResponse, error)
	UpdateUser(id string, user request.UpdateUserRequest) (response.UserResponse, error)
	DeleteUser(id string) (response.UserResponse, error)
}

type UserAuthService interface {
	Register(user request.RegisterUserRequest) (response.RegisterUserResponse, error)
	Login(users request.LoginUserRequest) (response.LoginUserResponse, error)
	Logout(users request.LogoutUserRequest) error
	Validate(token string) (response.LoginUserResponse, error)
}

type FileUploadService interface {
	UploadFile(file request.FileUploadRequest) (response.FileUploadResponse, error)
	GetAllFileUpload() ([]response.FileUploadResponse, error)

	GetFileUploadByFileID(id string) (response.FileUploadResponse, error)
	GetFileUploadByFileName(fileName string) (response.FileUploadResponse, error)
	GetFileUploadByFileHash(fileHash string) (response.FileUploadResponse, error)

	GetFilesUploadByAuthorID(authorID string) ([]response.FileUploadResponse, error)
	GetFilesUploadByFileName(fileName string) ([]response.FileUploadResponse, error)

	UpdateFileUpload(id string, fileUpdate request.FileUploadRequest) (response.FileUploadResponse, error)

	DeleteFileUploadByFileID(id string) (response.FileUploadResponse, error)
	DeleteFileUploadByFileName(fileName string) (response.FileUploadResponse, error)
	DeleteFileUploadByFileHash(fileHash string) (response.FileUploadResponse, error)
}
