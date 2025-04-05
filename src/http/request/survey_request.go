package request

import "time"

type SurveyorLinkRequest struct {
	ID         string `json:"id"`
	SurveyID   string `json:"survey_id"`
	SurveyorID string `json:"surveyor_id"`
}

type SurveyRequest struct {
	Title        string                `json:"title"`
	FormNumber   string                `json:"form_number"`
	QuestorName  string                `json:"questor_name"`
	FAT          string                `json:"fat"`
	CustomerName string                `json:"customer_name"`
	Address      string                `json:"address"`
	NodeFDT      string                `json:"node_fdt"`
	SurveyDate   time.Time             `json:"survey_date"`
	Surveyors    []SurveyorLinkRequest `json:"surveyors"`
	ImageID      string                `json:"image_id"`
}

type UpdateSurveyRequest struct {
	Title        *string               `json:"title,omitempty"`
	FormNumber   *string               `json:"form_number,omitempty"`
	QuestorName  *string               `json:"questor_name,omitempty"`
	FAT          *string               `json:"fat,omitempty"`
	CustomerName *string               `json:"customer_name,omitempty"`
	Address      *string               `json:"address,omitempty"`
	NodeFDT      *string               `json:"node_fdt,omitempty"`
	SurveyDate   *time.Time            `json:"survey_date,omitempty"`
	Surveyors    []SurveyorLinkRequest `json:"surveyors,omitempty"`
	ImageID      *string               `json:"image_id,omitempty"`
}

type SurveyReportRequest struct {
	SurveyID string  `json:"survey_id" validate:"required"`
	Status   string  `json:"status" validate:"required"`
	Remark   string  `json:"remark" validate:"required"`
	ImageID  *string `json:"image_id"`
}

type UpdateSurveyReportRequest struct {
	Remark  *string `json:"remark,omitempty"`
	Status  *string `json:"status,omitempty"`
	ImageID *string `json:"image_id,omitempty"`
}

type NotifyRequest struct {
	UserID        string `json:"user_id"`
	FileID        string `json:"file_id"`
	NotifyType    string `json:"notify_type"`
	NotifyStatus  string `json:"notify_status"`
	NotifyMessage string `json:"notify_message"`
	IsRead        bool   `json:"is_read"`
}
