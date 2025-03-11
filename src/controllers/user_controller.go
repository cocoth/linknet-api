package controllers

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/controllers/utils"
	"github.com/cocoth/linknet-api/src/http/request"
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

func (u *UserController) Create(c *gin.Context) {
	var createReq request.CreateUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.userService.Create(createReq)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.RespondWithError(c, http.StatusConflict, err.Error())
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, userRes)
}

func (u *UserController) Update(c *gin.Context) {
	var updateUserReq request.UpdateUserRequest

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	updateUserReq.Id = id

	if _, err := u.userService.Update(updateUserReq); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := u.userService.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondWithError(c, http.StatusNotFound, "not found")
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil)
}

func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.userService.GetAll()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetById(c *gin.Context) {
	id := c.Param("id")

	user, err := u.userService.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			utils.RespondWithError(c, http.StatusNotFound, "user not found")
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, user)
}
