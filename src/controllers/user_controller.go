package controllers

import (
	"net/http"

	ctrlUtils "github.com/cocoth/linknet-api/src/controllers/utils"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
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
	var createReq request.UserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(createReq.Password) < 8 {
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, "Password must be at least 8 characters long")
		return
	}

	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can create user!")
		return
	}

	userRes, err := u.userService.Create(createReq)
	if err != nil {
		if err.Error() == "user with that email already exists" {
			ctrlUtils.RespondWithError(c, http.StatusConflict, err.Error())
		} else {
			ctrlUtils.HandleGormError(c, err)
		}
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusCreated, userRes)
}

func (u *UserController) Update(c *gin.Context) {
	var updateUserReq request.UpdateUserRequest

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		user, err := u.userService.GetById(id)
		if err != nil {
			ctrlUtils.RespondWithError(c, http.StatusNotFound, "user not found")
			return
		}
		if user.Role.Name != "user" || (updateUserReq.Role != nil && updateUserReq.Role.Name != "user") {
			ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can update user!")
			return
		}
	}

	user, err := u.userService.Update(id, updateUserReq)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ctrlUtils.RespondWithSuccess(c, http.StatusOK, user)

}

func (u *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can delete user!")
		return
	}
	user, err := u.userService.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctrlUtils.RespondWithError(c, http.StatusNotFound, "not found")
		} else {
			ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, user)
}

func (u *UserController) GetAll(c *gin.Context) {
	var users []response.UserResponse
	var err error

	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can get user!")
		return
	}

	email := c.Query("email")
	name := c.Query("name")
	phone := c.Query("phone")
	role := c.Query("role")
	status := c.Query("status")
	contractor := c.Query("contractor")

	if email != "" {
		users, err = u.userService.GetUsersByEmail(email)
	} else if name != "" {
		users, err = u.userService.GetUsersByName(name)
	} else if phone != "" {
		users, err = u.userService.GetUsersByPhone(phone)
	} else if role != "" {
		users, err = u.userService.GetUsersByRole(role)
	} else if status != "" {
		users, err = u.userService.GetUsersByStatus(status)
	} else if contractor != "" {
		users, err = u.userService.GetUsersByContractor(contractor)
	} else {
		users, err = u.userService.GetAll()
	}

	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetById(c *gin.Context) {
	id := c.Param("id")
	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can get user!")
		return
	}

	user, err := u.userService.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctrlUtils.RespondWithError(c, http.StatusNotFound, "user not found")
		} else {
			ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, user)
}

func (u *UserController) GetAllRole(c *gin.Context) {
	roles, err := u.userService.GetAllRole()
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can create role!")
		return
	}
	ctrlUtils.RespondWithSuccess(c, http.StatusOK, roles)
}

func (u *UserController) CreateRole(c *gin.Context) {
	var createReq request.RoleRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		ctrlUtils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.Cookie("session_token")
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, err := u.userService.IsAdmin(token)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		ctrlUtils.RespondWithError(c, http.StatusUnauthorized, "only admin can create role!")
		return
	}

	roleRes, err := u.userService.CreateRole(createReq)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusCreated, roleRes)
}

func (u *UserController) GetUsersByName(c *gin.Context) {
	name := c.Query("name")
	users, err := u.userService.GetUsersByName(name)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetUsersByEmail(c *gin.Context) {
	email := c.Query("email")
	users, err := u.userService.GetUsersByEmail(email)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetUsersByPhone(c *gin.Context) {
	phone := c.Query("phone")
	users, err := u.userService.GetUsersByPhone(phone)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetUsersByRole(c *gin.Context) {
	role := c.Query("role")
	users, err := u.userService.GetUsersByRole(role)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetUsersByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := u.userService.GetUsersByStatus(status)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) GetUsersByContractor(c *gin.Context) {
	contractor := c.Query("contractor")
	users, err := u.userService.GetUsersByContractor(contractor)
	if err != nil {
		ctrlUtils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctrlUtils.RespondWithSuccess(c, http.StatusOK, users)
}
