package helper

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, response.WebResponse{
		Code:    code,
		Status:  "Error",
		Message: message,
		Data:    nil,
	})
}

func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, response.WebResponse{
		Code:    code,
		Status:  "Ok",
		Message: "",
		Data:    data,
	})
}

func HandleGormError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch err {
	case gorm.ErrRecordNotFound:
		RespondWithError(c, http.StatusNotFound, "record not found")
	case gorm.ErrDuplicatedKey:
		RespondWithError(c, http.StatusConflict, err.Error())
	default:
		RespondWithError(c, http.StatusInternalServerError, err.Error())
	}
}
