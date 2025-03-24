package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService services.FileUploadService
	userService services.UserService
}

func NewFileController(fileService services.FileUploadService, userService services.UserService) *FileController {
	return &FileController{
		fileService: fileService,
		userService: userService,
	}
}

func (f *FileController) UploadFile(c *gin.Context) {
	var fileUploadRequest request.FileUploadRequest

	file, err := c.FormFile("file")
	if err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}
	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}

	exp, userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	userRes, err := f.userService.GetUserById(userId)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if float64(time.Now().Unix()) > exp {
		helper.RespondWithError(c, http.StatusUnauthorized, "Token Expired")
		return
	}

	uploadDir := os.Getenv("APP_UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filePath := filepath.Join(uploadDir, file.Filename)

	// Check if file already exists on the file system
	if _, err := os.Stat(filePath); err == nil {
		helper.RespondWithError(c, 400, "File already exists on the server")
		return
	}

	fileBytes, _ := file.Open()
	defer fileBytes.Close()

	hash := utils.CalculateHashByBuffer(fileBytes)
	// Check if file already exists in the database
	existingFile, err := f.fileService.GetFileUploadByFileHash(hash)
	if err == nil && existingFile != (response.FileUploadResponse{}) {
		helper.RespondWithError(c, 400, "File already exists in the database")
		return
	}

	// Upload the file to specific dir
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}

	fileHash, err := utils.CalculateHash(filePath)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to calculate file hash")
		return
	}

	fileUploadRequest.FileName = file.Filename
	fileUploadRequest.FileUri = filePath
	fileUploadRequest.FileHash = fileHash
	fileUploadRequest.AuthorID = userRes.ID

	// Upload metadata to database
	fileRes, err := f.fileService.UploadFile(fileUploadRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, fileRes)

}

func (f *FileController) GetAllFileUpload(c *gin.Context) {
	var files []response.FileUploadResponse
	var err error

	qFileID := c.Query("id")
	qFileName := c.Query("filename")
	qfileHash := c.Query("filehash")
	qAuthorID := c.Query("authorid")

	if qFileID != "" {
		file, err := f.fileService.GetFileUploadByFileID(qFileID)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, file)
		return
	} else if qFileName != "" {
		file, err := f.fileService.GetFileUploadByFileName(qFileName)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, file)
		return
	} else if qfileHash != "" {
		file, err := f.fileService.GetFileUploadByFileHash(qfileHash)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, file)
		return
	} else if qAuthorID != "" {
		files, err = f.fileService.GetFilesUploadByAuthorID(qAuthorID)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "File not found")
			return
		}
	} else {
		files, err = f.fileService.GetAllFileUpload()

	}

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, files)
}

func (f *FileController) UpdateFileUpload(c *gin.Context) {
	var fileReq request.FileUploadRequest
	var file response.FileUploadResponse
	var err error

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
	if err == nil && existingFile != (response.FileUploadResponse{}) {
		helper.RespondWithError(c, 400, "File already exists in the database")
		return
	}

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}

	exp, userId, err := utils.ValidateJWTToken(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	userRes, err := f.userService.GetUserById(userId)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if float64(time.Now().Unix()) > exp {
		helper.RespondWithError(c, http.StatusUnauthorized, "Token Expired")
		return
	}

	uploadDir := os.Getenv("APP_UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}

	filePath := filepath.Join(uploadDir, Updatedfile.Filename)

	// Remove the old file from the file system
	err = os.Remove(file.FileUri)

	if err != nil {
		utils.Debug(file.FileUri + " =====> " + err.Error())
		helper.RespondWithError(c, 404, "no such file or directory")
		return
	}

	// Save the new file to the file system
	if err := c.SaveUploadedFile(Updatedfile, filePath); err != nil {
		helper.RespondWithError(c, 400, err.Error())
		return
	}

	fileHash, err := utils.CalculateHash(filePath)
	if err != nil {
		helper.RespondWithError(c, 500, "Failed to calculate file hash")
		return
	}

	fileReq.FileName = Updatedfile.Filename
	fileReq.FileUri = filePath
	fileReq.FileHash = fileHash
	fileReq.AuthorID = userRes.ID

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
	helper.RespondWithSuccess(c, http.StatusOK, file)
}
