package controllers

import (
	"net/http"
	"strconv"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/http/request"
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

	qID := c.Query("id")
	qUserID := c.Query("user_id")
	qFileID := c.Query("file_id")
	qNotifyMessage := c.Query("notify_message")
	qNotifyStatus := c.Query("notify_status")
	qNotifyType := c.Query("notify_type")
	qIsRead := c.Query("is_read")

	filters := map[string]interface{}{}

	if qID != "" {
		filters["id"] = qID
	}
	if qUserID != "" {
		filters["user_id"] = qUserID
	}
	if qFileID != "" {
		filters["file_id"] = qFileID
	}
	if qNotifyMessage != "" {
		filters["notify_message"] = qNotifyMessage
	}
	if qNotifyStatus != "" {
		filters["notify_status"] = qNotifyStatus
	}
	if qNotifyType != "" {
		filters["notify_type"] = qNotifyType
	}
	if qIsRead != "" {
		isRead, err := strconv.ParseBool(qIsRead)
		if err != nil {
			helper.RespondWithError(c, http.StatusBadRequest, "Invalid value for is_read, must be true or false")
			return
		}
		filters["is_read"] = isRead
	}

	notifies, err := n.notifyService.GetNotifyWithFilters(filters)

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(notifies) == 0 {
		helper.RespondWithError(c, http.StatusNotFound, "No Notify found with that given filters")
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
	SendNotification(notif.UserID, notif)
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
	SendNotification(notif.UserID, notif)
	helper.RespondWithSuccess(c, http.StatusOK, notif)
}

func (n *NotifyController) DeleteNotify(c *gin.Context) {
	notifyID := c.Param("id")
	if notifyID == "" {
		helper.RespondWithError(c, http.StatusBadRequest, "Notify ID is required")
		return
	}
	notif, err := n.notifyService.DeleteNotify(notifyID)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, notif)
}
