package controllers

import (
	"net/http"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
	"github.com/cocoth/linknet-api/src/http/response"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-gonic/gin"
)

type ISmartControler struct {
	ismartService services.ISmartService
}

func NewISmartController(ismartService services.ISmartService) *ISmartControler {
	return &ISmartControler{
		ismartService: ismartService,
	}
}
func (i *ISmartControler) GetAllISmart(c *gin.Context) {

	filters := map[string]interface{}{}

	qID := c.Query("id")
	qFiberNode := c.Query("fiber_node")
	qAddress := c.Query("address")
	qCoordinate := c.Query("coordinate")
	qStreet := c.Query("street")

	if qID != "" {
		filters["id"] = qID
	}
	if qFiberNode != "" {
		filters["fiber_node"] = qFiberNode
	}
	if qAddress != "" {
		filters["address"] = qAddress
	}
	if qCoordinate != "" {
		filters["coordinate"] = qCoordinate
	}
	if qStreet != "" {
		filters["street"] = qStreet
	}

	ismarts, err := i.ismartService.GetISmartsWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(ismarts) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No ISmart found with that given filters")
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, ismarts)
}

func (i *ISmartControler) CreateISmart(c *gin.Context) {
	var iSmartRequest request.ISmartRequest

	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "You are not authorized to create ISmart")
		return
	}

	if err := c.ShouldBindJSON(&iSmartRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	iSmartResponse, err := i.ismartService.CreateISmart(iSmartRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusCreated, iSmartResponse)
}
func (i *ISmartControler) UpdateISmart(c *gin.Context) {
	var iSmartRequest request.ISmartRequest
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "You are not authorized to update ISmart")
		return
	}
	if err := c.ShouldBindJSON(&iSmartRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	iSmartResponse, err := i.ismartService.UpdateISmart(id, iSmartRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, iSmartResponse)
}
func (i *ISmartControler) DeleteISmart(c *gin.Context) {
	token, exsist := c.Get("current_user")
	if !exsist {
		helper.RespondWithError(c, http.StatusUnauthorized, "No token provided")
		return
	}
	currentResUser := token.(response.UserResponse)

	if currentResUser.Role.Name != "admin" {
		helper.RespondWithError(c, http.StatusUnauthorized, "You are not authorized to delete ISmart")
		return
	}
	id := c.Param("id")
	iSmartResponse, err := i.ismartService.DeleteISmart(id)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, iSmartResponse)
}
