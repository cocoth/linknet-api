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

type SurveyRepo interface {
	GetAllSurvey() ([]models.Survey, error)
	GetSurveyByID(id string) (models.Survey, error)
	GetSurveyByTitle(title string) (models.Survey, error)
	GetSurveyByFormNumber(formNumber string) (models.Survey, error)
	GetSurveyByQuestorName(questorName string) (models.Survey, error)

	GetSurveyByFAT(fat string) (models.Survey, error)

	GetSurveyByCustomerName(customerName string) (models.Survey, error)
	GetSurveyByAddress(address string) (models.Survey, error)

	GetSurveyByNodeFDT(nodeFDT string) (models.Survey, error)
	GetSurveyBySurveyDate(surveyDate string) (models.Survey, error)
	GetSurveyBySurveyorID(surveyorID string) (models.Survey, error)

	GetSurveyByImageID(imageID string) (models.Survey, error)

	GetSurveysByTitle(title string) ([]models.Survey, error)
	GetSurveysByQuestorName(questorName string) ([]models.Survey, error)
	GetSurveysByCustomerName(customerName string) ([]models.Survey, error)
	GetSurveysByAddress(address string) ([]models.Survey, error)
	GetSurveysBySurveyorName(surveyorName string) ([]models.Survey, error)

	CreateSurvey(survey models.Survey) (models.Survey, error)
	UpdateSurvey(survey models.Survey) (models.Survey, error)
	DeleteSurvey(id string) (models.Survey, error)
}

type SurveyReportRepo interface {
	GetAllSurveyReport() ([]models.SurveyReport, error)
	GetSurveyReportByID(id string) (models.SurveyReport, error)
	GetSurveyReportBySurveyID(surveyID string) (models.SurveyReport, error)
	GetSurveyReportByRemark(remark string) (models.SurveyReport, error)
	GetSurveyReportByStatus(status string) (models.SurveyReport, error)

	CreateSurveyReport(surveyReport models.SurveyReport) (models.SurveyReport, error)
	UpdateSurveyReport(surveyReport models.SurveyReport) (models.SurveyReport, error)
	DeleteSurveyReport(id string) (models.SurveyReport, error)
}

type NotifyRepo interface {
	GetAllNotify() ([]models.Notify, error)
	GetNotifyByID(id string) (models.Notify, error)
	GetNotifyByUserID(userID string) (models.Notify, error)
	GetNotifyByFileID(fileID string) (models.Notify, error)
	GetNotifyByNotifyType(notifyType string) (models.Notify, error)
	GetNotifyByNotifyStatus(notifyStatus string) (models.Notify, error)
	GetNotifyByNotifyMessage(notifyMessage string) (models.Notify, error)

	CreateNotify(notify models.Notify) (models.Notify, error)
	UpdateNotify(notify models.Notify) (models.Notify, error)
	DeleteNotify(id string) (models.Notify, error)
}
