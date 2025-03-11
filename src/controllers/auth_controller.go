package controllers

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/controllers/utils"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	authService services.UserAuthService
}

func NewAuthController(service services.UserAuthService) *UserAuthController {
	return &UserAuthController{
		authService: service,
	}
}

func (u *UserAuthController) Register(c *gin.Context) {
	var createReq request.RegisterUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Register(createReq)
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

func (u *UserAuthController) Login(c *gin.Context) {
	var createReq request.LoginUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Login(createReq)
	if err != nil {
		if err.Error() == "invalid credentials" {
			utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, userRes)
}

func (u *UserAuthController) Logout() {
	// TODO
}

func (u *UserAuthController) RefreshToken() {
	// TODO
}
