package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	authService services.UserAuthService
	userService services.UserService
}

func NewAuthController(service services.UserAuthService, userService services.UserService) *UserAuthController {
	return &UserAuthController{
		authService: service,
		userService: userService,
	}
}

func (u *UserAuthController) Register(c *gin.Context) {
	var createReq request.RegisterUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Register(createReq)
	if err != nil {
		if err.Error() == "user with that email already exists" {
			helper.RespondWithError(c, http.StatusConflict, err.Error())
		} else if err.Error() == "user with that phone already exists" {
			helper.RespondWithError(c, http.StatusConflict, err.Error())
		} else {
			helper.HandleGormError(c, err)
		}
		return
	}
	helper.RespondWithSuccess(c, http.StatusCreated, userRes)
}

func (u *UserAuthController) Login(c *gin.Context) {
	var createReq request.LoginUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Login(createReq)
	if err != nil {
		if err.Error() == "invalid credentials" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else if err.Error() == "record not found" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else {
			helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	csrfToken := userRes.CsrfToken
	sessionToken := userRes.SessionToken
	domain := os.Getenv("APP_DOMAIN")

	c.SetCookie("csrf_token", csrfToken, int((24 * time.Hour).Seconds()), "/", domain, false, false)
	c.SetCookie("session_token", sessionToken, int((24 * time.Hour).Seconds()), "/", domain, false, true)
	helper.RespondWithSuccess(c, http.StatusOK, userRes)
}

func (u *UserAuthController) Logout(c *gin.Context) {
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, "No session token provided")
		return
	}

	if sessionToken == "" {
		helper.RespondWithError(c, http.StatusBadRequest, "No session token provided")
		return
	}

	err = u.authService.Logout(sessionToken)
	if err != nil {
		if err.Error() == "record not found" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else {
			helper.RespondWithError(c, http.StatusInternalServerError, "Failed to log out")
		}
		return
	}

	domain := os.Getenv("APP_DOMAIN")
	c.SetCookie("csrf_token", "", -1, "/", domain, false, false)
	c.SetCookie("session_token", "", -1, "/", domain, false, true)

	helper.RespondWithSuccess(c, http.StatusOK, "Logged out successfully")
}
func (u *UserAuthController) Validate(c *gin.Context) {
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}

	currentResUser := token.(response.UserResponse)

	helper.RespondWithSuccess(c, http.StatusOK, currentResUser)
}
func (u *UserAuthController) ValidateAdmin(c *gin.Context) {
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}

	currentResUser := token.(response.UserResponse)
	fmt.Println("currentResUser: ", currentResUser.Role.Name)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusForbidden, "only admin can access this resource!")
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, currentResUser)
}

func (u *UserAuthController) CheckIsAdmin(c *gin.Context) {
	var createReq request.LoginUserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	userRes, err := u.authService.Login(createReq)
	if err != nil {
		if err.Error() == "invalid credentials" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else if err.Error() == "record not found" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		} else {
			helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	isadmin, _, err := u.userService.IsAdmin(userRes.SessionToken)

	if err != nil {
		helper.RespondWithError(c, http.StatusForbidden, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusForbidden, "You are not authorized to access this resource")
		return
	}

	sessionToken := userRes.SessionToken
	domain := os.Getenv("APP_DOMAIN")

	c.SetCookie("session_token", sessionToken, int((24 * time.Hour).Seconds()), "/", domain, false, true)
	helper.RespondWithSuccess(c, http.StatusOK, userRes)
}
