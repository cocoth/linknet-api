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

	IsAdmin(token string) (status bool, userResponse response.UserResponse, err error)
	CheckToken(token string) (status bool, userResponse response.UserResponse, err error)

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

type SurveyService interface {
	GetAllSurvey() ([]response.SurveyResponse, error)
	GetSurveyByID(id string) (response.SurveyResponse, error)
	GetSurveyByTitle(title string) (response.SurveyResponse, error)
	GetSurveyByFormNumber(formNumber string) (response.SurveyResponse, error)
	GetSurveyByQuestorName(questorName string) (response.SurveyResponse, error)
	GetSurveyByFAT(fat string) (response.SurveyResponse, error)
	GetSurveyByCustomerName(customerName string) (response.SurveyResponse, error)
	GetSurveyByAddress(address string) (response.SurveyResponse, error)
	GetSurveyByNodeFDT(nodeFDT string) (response.SurveyResponse, error)
	GetSurveyBySurveyDate(surveyDate string) (response.SurveyResponse, error)
	GetSurveyBySurveyorID(surveyorID string) (response.SurveyResponse, error)

	GetSurveyByImageID(imageID string) (response.SurveyResponse, error)

	GetSurveysByTitle(title string) ([]response.SurveyResponse, error)
	GetSurveysByQuestorName(questorName string) ([]response.SurveyResponse, error)
	GetSurveysByCustomerName(customerName string) ([]response.SurveyResponse, error)
	GetSurveysByAddress(address string) ([]response.SurveyResponse, error)
	GetSurveysBySurveyorName(surveyorName string) ([]response.SurveyResponse, error)

	CreateSurvey(survey request.SurveyRequest) (response.SurveyResponse, error)
	UpdateSurvey(id string, survey request.UpdateSurveyRequest) (response.SurveyResponse, error)
	DeleteSurvey(id string) (response.SurveyResponse, error)
}

type SurveyReportService interface {
	GetAllSurveyReport() ([]response.SurveyReportResponse, error)
	GetSurveyReportByID(id string) (response.SurveyReportResponse, error)
	GetSurveyBySurveyID(surveyID string) (response.SurveyReportResponse, error)
	GetSurveyReportByStatus(status string) (response.SurveyReportResponse, error)
	GetSurveyReportByRemark(remark string) (response.SurveyReportResponse, error)

	CreateSurveyReport(surveyReport request.SurveyReportRequest) (response.SurveyReportResponse, error)
	UpdateSurveyReport(id string, surveyReport request.UpdateSurveyReportRequest) (response.SurveyReportResponse, error)
	DeleteSurveyReport(id string) (response.SurveyReportResponse, error)
}

type NotifyService interface {
	GetAllNotify() ([]response.NotifyResponse, error)
	GetNotifyByID(id string) (response.NotifyResponse, error)
	GetNotifyByUserID(userID string) (response.NotifyResponse, error)
	GetNotifyByFileID(fileID string) (response.NotifyResponse, error)
	GetNotifyByNotifyType(notifyType string) (response.NotifyResponse, error)
	GetNotifyByNotifyStatus(notifyStatus string) (response.NotifyResponse, error)
	GetNotifyByNotifyMessage(notifyMessage string) (response.NotifyResponse, error)

	CreateNotify(notify request.NotifyRequest) (response.NotifyResponse, error)
	UpdateNotify(id string, notify request.NotifyRequest) (response.NotifyResponse, error)
	DeleteNotify(id string) (response.NotifyResponse, error)
}
