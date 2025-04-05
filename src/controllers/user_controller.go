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

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create user!")
		return
	}

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(createReq.Password) < 8 {
		helper.RespondWithError(c, http.StatusBadRequest, "Password must be at least 8 characters long")
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

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	if currentResUser.Role.Name != "admin" {
		// If not admin, ensure the user can only update their own data
		if currentResUser.ID != id {
			helper.RespondWithError(c, http.StatusUnauthorized, "You can only update your own data")
			return
		}

		// Restrict role updates for non-admin users
		if updateUserReq.Role != nil && updateUserReq.Role.Name != "user" {
			helper.RespondWithError(c, http.StatusUnauthorized, "Only admin can update roles")
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
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
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
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	qID := c.Query("id")
	qName := c.Query("name")
	qEmail := c.Query("email")
	qPhone := c.Query("phone")
	qRole := c.Query("role")
	qStatus := c.Query("status")
	qContractor := c.Query("contractor")
	qCallSign := c.Query("callsign")

	filters := map[string]interface{}{}

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithSuccess(c, http.StatusOK, currentResUser)
		return
	} else if currentResUser.Role.Name == "admin" {
		if qID != "" {
			filters["id"] = qID
		}
		if qName != "" {
			filters["name"] = qName
		}
		if qEmail != "" {
			filters["email"] = qEmail
		}
		if qPhone != "" {
			filters["phone"] = qPhone
		}
		if qRole != "" {
			filters["role"] = qRole
		}
		if qStatus != "" {
			filters["status"] = qStatus
		}
		if qContractor != "" {
			filters["contractor"] = qContractor
		}
		if qCallSign != "" {
			filters["callsign"] = qCallSign
		}
	}

	users, err := u.userService.GetUsersWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(users) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No users found with that given filters")
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, users)
}

func (u *UserController) CreateRole(c *gin.Context) {
	var createReq request.RoleRequest

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can create role!")
		return
	}

	if err := c.ShouldBindJSON(&createReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
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

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can get role!")
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

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can update role!")
		return
	}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

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
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "only admin can delete role!")
		return
	}
	id := c.Param("id")

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
