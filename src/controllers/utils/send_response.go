package utils

import (
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/gin-gonic/gin"
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
