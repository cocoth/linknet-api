package controllers

import (
	"net/http"
	"os"
	"time"

	ctrlUtils "github.com/cocoth/linknet-api/src/controllers/utils"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
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
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Register(createReq)
	if err != nil {
		if err.Error() == "email already exists" {
			ctrlUtils.RespondWithError(c, http.StatusConflict, err.Error())
		} else {
			ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctrlUtils.RespondWithSuccess(c, http.StatusCreated, userRes)
}

func (u *UserAuthController) Login(c *gin.Context) {
	var createReq request.LoginUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Login(createReq)
	if err != nil {
		if err.Error() == "invalid credentials" {
			ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else if err.Error() == "record not found" {
			ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else {
			ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	csrfToken := userRes.CsrfToken
	sessionToken := userRes.SessionToken
	domain := os.Getenv("APP_DOMAIN")

	c.SetCookie("csrf_token", csrfToken, int((24 * time.Hour).Seconds()), "/", domain, false, false)
	c.SetCookie("session_token", sessionToken, int((24 * time.Hour).Seconds()), "/", domain, false, true)
	ctrlUtils.RespondWithSuccess(c, http.StatusOK, userRes)
}

func (u *UserAuthController) Logout(c *gin.Context) {
	domain := os.Getenv("APP_DOMAIN")
	c.SetCookie("csrf_token", "", -1, "/", domain, false, false)
	c.SetCookie("session_token", "", -1, "/", domain, false, true)

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, "Logged out")
}
func (u *UserAuthController) Validate(c *gin.Context) {
	user, exsist := c.Get("user")

	if !exsist {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	userRes, err := u.authService.Validate(user.(response.UserResponse).Id)

	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, userRes)
}
