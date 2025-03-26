package controllers

import (
	"net/http"
	"strconv"

	"github.com/cocoth/linknet-api/src/controllers/helper"
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

func (u *UserController) CreateUser(c *gin.Context) {
	var createReq request.UserRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(createReq.Password) < 8 {
		helper.RespondWithError(c, http.StatusBadRequest, "Password must be at least 8 characters long")
		return
	}

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create user!")
		return
	}

	userRes, err := u.userService.CreateUser(createReq)
	if err != nil {
		if err.Error() == "user with that email already exists" {
			helper.RespondWithError(c, http.StatusConflict, err.Error())
		} else {
			helper.HandleGormError(c, err)
		}
		return
	}

	helper.RespondWithSuccess(c, http.StatusCreated, userRes)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var updateUserReq request.UpdateUserRequest

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		user, err := u.userService.GetUserById(id)
		if err != nil {
			helper.RespondWithError(c, http.StatusNotFound, "user not found")
			return
		}
		if user.Role.Name != "user" || (updateUserReq.Role != nil && updateUserReq.Role.Name != "user") {
			helper.RespondWithError(c, http.StatusUnauthorized, "only admin can update user!")
			return
		}
	}

	user, err := u.userService.UpdateUser(id, updateUserReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can delete user!")
		return
	}
	_, errDelete := u.userService.DeleteUser(id)
	if errDelete != nil {
		if errDelete.Error() == "record not found" {
			helper.RespondWithError(c, http.StatusNotFound, "not found")
		} else {
			helper.RespondWithError(c, http.StatusInternalServerError, errDelete.Error())
		}
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "user deleted successfully")
}

func (u *UserController) GetAll(c *gin.Context) {
	var users []response.UserResponse
	var err error

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	_, _, errIsAdmin := u.userService.CheckToken(token)
	if errIsAdmin != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, errIsAdmin.Error())
		return
	}

	qID := c.Query("id")
	qName := c.Query("name")
	qEmail := c.Query("email")
	qPhone := c.Query("phone")
	qRole := c.Query("role")
	qContractor := c.Query("contractor")
	qStatus := c.Query("status")

	if qID != "" {
		user, errUser := u.userService.GetUserById(qID)
		if errUser != nil {
			if errUser.Error() == "record not found" {
				helper.RespondWithError(c, http.StatusNotFound, "user not found")
			} else {
				helper.RespondWithError(c, http.StatusInternalServerError, errUser.Error())
			}
			return
		}
		helper.RespondWithSuccess(c, http.StatusOK, user)
		return
	} else if qName != "" {
		users, err = u.userService.GetUsersByName(qName)
	} else if qEmail != "" {
		users, err = u.userService.GetUsersByEmail(qEmail)
	} else if qPhone != "" {
		users, err = u.userService.GetUsersByPhone(qPhone)
	} else if qRole != "" {
		users, err = u.userService.GetUsersByRole(qRole)
	} else if qContractor != "" {
		users, err = u.userService.GetUsersByContractor(qContractor)
	} else if qStatus != "" {
		users, err = u.userService.GetUsersByStatus(qStatus)
	} else {
		users, err = u.userService.GetAll()
	}

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) CreateRole(c *gin.Context) {
	var createReq request.RoleRequest

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create role!")
		return
	}

	roleRes, err := u.userService.CreateRole(createReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusCreated, roleRes)
}

func (u *UserController) GetAllRole(c *gin.Context) {
	var roles []response.RoleResponse
	var err error

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create role!")
		return
	}

	qRoleID := c.Query("id")
	qRoleName := c.Query("name")

	if qRoleID != "" {
		roleID, err := strconv.ParseUint(qRoleID, 10, 32)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "Invalid role ID")
			return
		}
		role, errRoleID := u.userService.GetRoleByRoleID(uint(roleID))
		if errRoleID != nil {
			helper.RespondWithError(c, http.StatusNotFound, errRoleID.Error())
			return
		}

		helper.RespondWithSuccess(c, http.StatusOK, role)
		return
	} else if qRoleName != "" {
		role, errRoleName := u.userService.GetRoleByRoleName(qRoleName)
		if errRoleName != nil {
			helper.RespondWithError(c, http.StatusNotFound, errRoleName.Error())
			return
		}

		helper.RespondWithSuccess(c, http.StatusOK, role)
		return
	} else {
		roles, err = u.userService.GetAllRole()
		if err != nil {
			helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
			return
		}

	}

	helper.RespondWithSuccess(c, http.StatusOK, roles)
}

func (u *UserController) UpdateRole(c *gin.Context) {
	var updateReq request.RoleRequest

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can update role!")
		return
	}

	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	role, err := u.userService.UpdateRole(uint(roleID), updateReq)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, role)
}

func (u *UserController) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	token, err := c.Cookie("session_token")
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	isadmin, _, err := u.userService.IsAdmin(token)
	if err != nil {
		helper.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !isadmin {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can delete role!")
		return
	}
	roleID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	_, errDelete := u.userService.DeleteRoleByID(uint(roleID))
	if errDelete != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, errDelete.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, "role deleted successfully")
}
