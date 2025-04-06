package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService     services.FileUploadService
	filePermService services.FileUploadPermService
	userService     services.UserService
}

func NewFileController(fileService services.FileUploadService, filePermService services.FileUploadPermService, userService services.UserService) *FileController {
	return &FileController{
		fileService:     fileService,
		filePermService: filePermService,
		userService:     userService,
	}
}

func (f *FileController) UploadFile(c *gin.Context) {
	var fileUploadRequest request.FileUploadRequest

	file, err := c.FormFile("file")
	if err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}

	currentResUser := token.(response.UserResponse)

	fileBytes, _ := file.Open()
	defer fileBytes.Close()

	hash := utils.CalculateHashByBuffer(fileBytes)
	// Check if file already exists in the database
	existingFile, err := f.fileService.GetFileUploadByFileHash(hash)
	if err == nil && existingFile != (response.FileUploadResponse{}) {
		helper.RespondWithError(c, 400, "File already exists in the database")
		return
	}

	fileBytes.Seek(0, io.SeekStart)

	uploadDir := os.Getenv("APP_UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filePath, err := utils.SaveMultipartFile(file)
	if err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}

	fileHash, err := utils.CalculateHash(filePath)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to calculate file hash")
		return
	}

	fileBytes.Seek(0, io.SeekStart)

	buffer := make([]byte, 512)
	_, err = fileBytes.Read(buffer)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to read file content")
		return
	}

	fileUploadRequest.FileName = file.Filename
	fileUploadRequest.FileType = http.DetectContentType(buffer)
	fileUploadRequest.FileUri = filePath
	fileUploadRequest.FileHash = fileHash
	fileUploadRequest.AuthorID = currentResUser.ID

	// Upload metadata to database
	fileRes, err := f.fileService.UploadFile(fileUploadRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, fileRes)

}

func (f *FileController) GetAllFileUpload(c *gin.Context) {

	qFileID := c.Query("id")
	qFileName := c.Query("filename")
	qfileHash := c.Query("filehash")
	qAuthorID := c.Query("authorid")

	filters := map[string]interface{}{}

	if qFileID != "" {
		filters["id"] = qFileID
	}
	if qFileName != "" {
		filters["filename"] = qFileName
	}
	if qfileHash != "" {
		filters["filehash"] = qfileHash
	}
	if qAuthorID != "" {
		filters["authorid"] = qAuthorID
	}

	files, err := f.fileService.GetFilesWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(files) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No Files found with that given filters")
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, files)
}

func (f *FileController) UpdateFileUpload(c *gin.Context) {
	var fileReq request.FileUploadRequest
	var file response.FileUploadResponse
	var err error

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.ID != fileReq.AuthorID {
		if currentResUser.Role.Name != "admin" {
			helper.RespondWithError(c, http.StatusUnauthorized, "only admin or the file owner can update the file!")
			return
		}
	}

	qFileID := c.Param("id")

	file, err = f.fileService.GetFileUploadByFileID(qFileID)
	if err != nil {
		helper.RespondWithError(c, http.StatusNotFound, "File not found")
		return
	}

	Updatedfile, err := c.FormFile("file")
	if err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}

	fileBytes, _ := Updatedfile.Open()
	defer fileBytes.Close()

	hash := utils.CalculateHashByBuffer(fileBytes)
	// Check if file already exists in the database
	existingFile, err := f.fileService.GetFileUploadByFileHash(hash)
	if err == nil && existingFile != (response.FileUploadResponse{}) && existingFile.ID != file.ID {
		helper.RespondWithError(c, 400, "File already exists in the database")
		return
	}

	// Remove the old file from the file system
	err = os.Remove(file.FileUri)

	if err != nil {
		helper.RespondWithError(c, 404, "no such file or directory")
		return
	}

	uploadDir := os.Getenv("APP_UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	filePath, err := utils.SaveMultipartFile(Updatedfile)
	if err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}

	fileHash, err := utils.CalculateHash(filePath)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to calculate file hash")
		return
	}

	fileBytes.Seek(0, io.SeekStart)

	buffer := make([]byte, 512)
	_, err = fileBytes.Read(buffer)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to read file content")
		return
	}

	fileReq.FileName = Updatedfile.Filename
	fileReq.FileType = http.DetectContentType(buffer)
	fileReq.FileUri = filePath
	fileReq.FileHash = fileHash
	fileReq.AuthorID = currentResUser.ID

	file, err = f.fileService.UpdateFileUpload(file.ID, fileReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, file)
}

func (f *FileController) DownloadFile(c *gin.Context) {
	var file response.FileUploadResponse
	var err error

	qFileID := c.Query("id")
	qFileName := c.Query("filename")
	qfileHash := c.Query("filehash")

	if qFileID != "" {
		file, err = f.fileService.GetFileUploadByFileID(qFileID)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else if qFileName != "" {
		file, err = f.fileService.GetFileUploadByFileName(qFileName)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else if qfileHash != "" {
		file, err = f.fileService.GetFileUploadByFileHash(qfileHash)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else {
		helper.RespondWithError(c, http.StatusBadRequest, "Invalid query parameter")
		return
	}
	if _, err := os.Stat(file.FileUri); os.IsNotExist(err) {
		helper.RespondWithError(c, http.StatusNotFound, "File not found on server")
		return
	}

	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	// c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.FileName))
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", file.FileName))
	c.Header("Content-Type", file.FileType)
	c.File(file.FileUri)

}

func (f *FileController) DeleteFileUpload(c *gin.Context) {
	var file response.FileUploadResponse
	var err error

	qFileID := c.Query("id")
	qFileName := c.Query("filename")
	qfileHash := c.Query("filehash")

	if qFileID != "" {
		file, err = f.fileService.DeleteFileUploadByFileID(qFileID)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else if qFileName != "" {
		file, err = f.fileService.DeleteFileUploadByFileName(qFileName)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else if qfileHash != "" {
		file, err = f.fileService.DeleteFileUploadByFileHash(qfileHash)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else {
		helper.RespondWithError(c, http.StatusBadRequest, "Invalid query parameter")
		return
	}

	err = os.Remove(file.FileUri)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "File deleted successfully")
}

func (f *FileController) RequestAccess(c *gin.Context) {
	var filePermission request.FilePermRequest

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, errPerm := f.userService.IsAdmin(token)
	if errPerm != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errPerm.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can access this!")
		return
	}

	if errBinding := c.ShouldBindJSON(&filePermission); errBinding != nil {
		helper.RespondWithError(c, http.StatusBadRequest, errBinding.Error())
		return
	}

	errReq := f.filePermService.RequestAccess(filePermission)
	if errReq != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, errReq.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "Access requested successfully")
}

func (f *FileController) ApproveAccess(c *gin.Context) {
	var filePermission request.FilePermRequest

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, errPerm := f.userService.IsAdmin(token)
	if errPerm != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errPerm.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can access this!")
		return
	}

	if errBinding := c.ShouldBindJSON(&filePermission); errBinding != nil {
		helper.RespondWithError(c, http.StatusBadRequest, errBinding.Error())
		return
	}

	errReq := f.filePermService.ApproveFileAccess(filePermission)
	if errReq != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, errReq.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "Access Approved")
}

func (f *FileController) RejectAccess(c *gin.Context) {
	var filePermission request.FilePermRequest

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := f.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can access this!")
		return
	}

	if err := c.ShouldBindJSON(&filePermission); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	errReq := f.filePermService.RejectFileAccess(filePermission)
	if errReq != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, errReq.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "Access Rejected")
}
