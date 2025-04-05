package response

import "time"

type SurveyorLinkResponse struct {
	ID         string     `json:"id"`
	SurveyID   string     `json:"survey_id"`
	SurveyorID string     `json:"surveyor_id"`
	CreatedAt  time.Time  `json:"CreatedAt"`
	UpdatedAt  time.Time  `json:"UpdatedAt"`
	DeletedAt  *time.Time `json:"DeletedAt"`
}

type SurveyResponse struct {
	ID           string                 `json:"id"`
	Title        string                 `json:"title"`
	FormNumber   string                 `json:"form_number"`
	QuestorName  string                 `json:"questor_name"`
	FAT          string                 `json:"fat"`
	CustomerName string                 `json:"customer_name"`
	Address      string                 `json:"address"`
	NodeFDT      string                 `json:"node_fdt"`
	SurveyDate   time.Time              `json:"survey_date"`
	Surveyors    []SurveyorLinkResponse `json:"surveyors"`
	ImageID      *string                `json:"image_id"`
	CreatedAt    time.Time              `json:"CreatedAt"`
	UpdatedAt    time.Time              `json:"UpdatedAt"`
	DeletedAt    *time.Time             `json:"DeletedAt"`
}

type SurveyReportResponse struct {
	ID        string     `json:"id"`
	SurveyID  string     `json:"survey_id"`
	Status    string     `json:"status"`
	Remark    string     `json:"remark"`
	ImageID   *string    `json:"image_id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}

type NotifyResponse struct {
	ID            string     `json:"id"`
	UserID        string     `json:"user_id"`
	FileID        string     `json:"file_id"`
	NotifyType    string     `json:"notify_type"`
	NotifyStatus  string     `json:"notify_status"`
	NotifyMessage string     `json:"notify_message"`
	IsRead        bool       `json:"is_read"`
	CreatedAt     time.Time  `json:"CreatedAt"`
	UpdatedAt     time.Time  `json:"UpdatedAt"`
	DeletedAt     *time.Time `json:"DeletedAt"`
}
