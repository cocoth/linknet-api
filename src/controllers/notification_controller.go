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

type NotifyController struct {
	notifyService services.NotifyService
}

func NewNotifyController(notifyService services.NotifyService) *NotifyController {
	return &NotifyController{notifyService: notifyService}
}

func (n *NotifyController) GetAllNotify(c *gin.Context) {
	var notifies []response.NotifyResponse
	var err error

	qUserID := c.Query("user_id")
	qFileID := c.Query("file_id")
	qNotifyMessage := c.Query("notify_message")
	qNotifyStatus := c.Query("notify_status")
	qNotifyType := c.Query("notify_type")
	qIsRead := c.Query("is_read")

	if qUserID != "" {
		notifies, err = n.notifyService.GetNotifyByUserID(qUserID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qFileID != "" {
		notifies, err = n.notifyService.GetNotifyByFileID(qFileID)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qNotifyMessage != "" {
		notifies, err = n.notifyService.GetNotifyByNotifyMessage(qNotifyMessage)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qNotifyStatus != "" {
		notifies, err = n.notifyService.GetNotifyByNotifyStatus(qNotifyStatus)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qNotifyType != "" {
		notifies, err = n.notifyService.GetNotifyByNotifyType(qNotifyType)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if qIsRead != "" {
		isRead, err := strconv.ParseBool(qIsRead)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "Invalid value for is_read, must be true or false")
			return
		}
		notifies, err = n.notifyService.GetNotifyByIsRead(isRead)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		notifies, err = n.notifyService.GetAllNotify()
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, notifies)
}

func (n *NotifyController) CreateNotify(c *gin.Context) {
	var notifyRequest request.NotifyRequest
	if err := c.ShouldBindJSON(&notifyRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	notif, err := n.notifyService.CreateNotify(notifyRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendNotification(notif)
	helper.RespondWithSuccess(c, http.StatusCreated, notif)
}

func (n *NotifyController) UpdateNotify(c *gin.Context) {
	var notifyRequest request.NotifyRequest
	if err := c.ShouldBindJSON(&notifyRequest); err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	notif, err := n.notifyService.UpdateNotify(notifyRequest)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	SendNotification(notif)
	helper.RespondWithSuccess(c, http.StatusOK, notif)
}

func (n *NotifyController) DeleteNotify(c *gin.Context) {
	notifyID := c.Param("id")
	_, err := n.notifyService.DeleteNotify(notifyID)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusNoContent, nil)
}
