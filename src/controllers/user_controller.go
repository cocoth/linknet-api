package controllers

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/data/request"
	"github.com/cocoth/linknet-api/src/data/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (u *UserController) respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, response.WebResponse{
		Code:    code,
		Status:  "Error",
		Message: message,
		Data:    nil,
	})
}

func (u *UserController) respondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, response.WebResponse{
		Code:    code,
		Status:  "Ok",
		Message: "",
		Data:    data,
	})
}

func (u *UserController) Create(c *gin.Context) {
	var createReq request.CreateUserReq

	if err := c.ShouldBindJSON(&createReq); err != nil {
		u.respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.userService.Create(createReq)
	if err != nil {
		if err.Error() == "Email already exists" {
			u.respondWithError(c, http.StatusBadRequest, err.Error())
			return
		} else {
			u.respondWithError(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	u.respondWithSuccess(c, http.StatusOK, userRes)
}

func (u *UserController) Update(c *gin.Context) {
	var updateUserReq request.UpdateUserReq

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		u.respondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	updateUserReq.Id = id

	if _, err := u.userService.Update(updateUserReq); err != nil {
		u.respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.respondWithSuccess(c, http.StatusOK, nil)
}

func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := u.userService.Delete(id)
	if err != nil {
		u.respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.respondWithSuccess(c, http.StatusOK, nil)
}

func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.userService.GetAll()
	if err != nil {
		u.respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.respondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetById(c *gin.Context) {
	id := c.Param("id")

	user, err := u.userService.GetById(id)
	if err != nil {
		u.respondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	u.respondWithSuccess(c, http.StatusOK, user)
}
